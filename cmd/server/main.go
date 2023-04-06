package main

import (
	"database/sql"
	"log"

	"time"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	handlers "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/handlers"
	db "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/repository/sqlc"
	usecase "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/usecase"
	config "github.com/pedrovitorrs/gofinances-backend/pkg/config"
	httpMiddleware "github.com/pedrovitorrs/gofinances-backend/pkg/web/middlewares"
)

func main() {
	config, err := config.LoadConfig(`config.json`)
	if err != nil {
		panic(err)
	}

	if config.Debug {
		log.Println("Service RUN on DEBUG mode")
	}

	conn, err := sql.Open(config.Database.Driver, config.Database.Source)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)

	e := echo.New()
	middL := httpMiddleware.InitMiddleware()
	e.Use(middL.CORS)

	timeoutContext := time.Duration(config.Context.Timeout) * time.Second

	userUseCase := usecase.NewUserUseCase(store, timeoutContext)
	handlers.NewUserHandler(e, userUseCase)

	log.Fatal(e.Start(config.Server.Address))
}
