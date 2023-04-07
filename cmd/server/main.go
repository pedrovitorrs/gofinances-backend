package main

import (
	"database/sql"
	"log"

	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	_ "github.com/pedrovitorrs/gofinances-backend/docs"
	handlers "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/handlers"
	db "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/repository/sqlc"
	usecase "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/usecase"
	config "github.com/pedrovitorrs/gofinances-backend/pkg/config"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Finances server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /api/v1
func main() {
	config, err := config.LoadConfig(`config.json`)
	if err != nil {
		panic(err)
	}

	conn, err := sql.Open(config.Database.Driver, config.Database.Source)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)

	e := echo.New()
	// middL := httpMiddleware.InitMiddleware()
	e.Use(middleware.CORS())
	if config.Debug {
		log.Println("Service RUN on DEBUG mode")
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	timeoutContext := time.Duration(config.Context.Timeout) * time.Second

	userUseCase := usecase.NewUserUseCase(store, timeoutContext)
	handlers.NewUserHandler(e, userUseCase)

	log.Fatal(e.Start(config.Server.Address))
}
