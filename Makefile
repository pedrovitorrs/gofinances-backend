createdb:
	createdb --username=postgres --owner=postgres go_finance

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:14-alpine

migrateup:
	migrate -path pkg/database/migrations -database "postgresql://postgres:password@localhost:5432/go_finance?sslmode=disable" -verbose up

migrationdrop:
	migrate -path pkg/database/migrations -database "postgresql://postgres:password@localhost:5432/go_finance?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

server:
	go run cmd/server/main.go

sqlc-gen:
	docker run --rm -v $$(pwd):/src -w /src kjconroy/sqlc generate

.PHONY: createdb postgres dropdb migrateup migrationdrop test server sqlc-gen