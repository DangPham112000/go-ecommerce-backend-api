GOOSE_DBSTRING = "root:root1234@tcp(127.0.0.1:3306)/shopgolang"
GOOSE_MIGRATION_DIR ?= sql/schema
# name app
APP_NAME := server

dev: 
	go run ./cmd/$(APP_NAME)
run:
	sudo docker compose up -d && go run ./cmd/$(APP_NAME)
kill:
	sudo docker compose kill 
up: 
	sudo docker compose up -d 
down:
	sudo docker compose down

upse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up
downse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down
resetse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset
	

.PHONY: run downse upse resetse