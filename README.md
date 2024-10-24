# MVC architecture

Users -> Controller -> Service -> Repo -> Models -> DBs

# SQL flow

## Migration

1. `make create_migration`
    -> goose will create file in sql/schema folder
2. Past table schema into created file
3. `make up_by_one`
    -> goose will create table for us in mysql -> shopgolang

## Create funcs let Go interact with Mysql

1. Create query file satisfy sqlc syntax in sql/queries folder
2. `make sqlgen`
    -> sqlc will create files in `internal/database` 
    -> funcs base on our query files will auto gen also


# Testing

```cmd
go test -v
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
open coverage.html
```

# Run

Run all necessary services (mysql, redis, kafka, kafka-ui) before run app: `make dcup`

Run app: `make dev`