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
	httpMiddleware "github.com/pedrovitorrs/gofinances-backend/pkg/web/middlewares"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbDriver := viper.GetString(`database.driver`)
	dbSource := viper.GetString(`database.source`)

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)

	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	middL := httpMiddleware.InitMiddleware()
	e.Use(middL.CORS)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userUseCase := usecase.NewUserUseCase(store, timeoutContext)
	handlers.NewUserHandler(e, userUseCase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
