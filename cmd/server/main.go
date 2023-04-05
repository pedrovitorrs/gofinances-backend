package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/labstack/echo"
	db "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/repository/sqlc"
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

	db.NewStore(conn)

	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	middL := httpMiddleware.InitMiddleware()
	e.Use(middL.CORS)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
