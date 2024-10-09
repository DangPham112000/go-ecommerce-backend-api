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

.PHONY: run