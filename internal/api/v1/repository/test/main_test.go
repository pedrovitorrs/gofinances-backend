package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	db "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/repository/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:password@localhost:5432/go_finance?sslmode=disable"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = db.New(conn)
	os.Exit(m.Run())
}
