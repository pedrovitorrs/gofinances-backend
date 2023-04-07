GOPATH := $(shell go env GOPATH)

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

doc-gen:
	$(GOPATH)/bin/swag init -g /cmd/server/main.go

server:
	go run cmd/server/main.go

sqlc-gen:
	docker run --rm -v $$(pwd):/src -w /src kjconroy/sqlc generate

.PHONY: createdb postgres dropdb migrateup migrationdrop test doc-gen server sqlc-gen