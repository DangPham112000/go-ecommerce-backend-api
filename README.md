# MVC architecture

Users -> Controller -> Service -> Repo -> Models -> DBs

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