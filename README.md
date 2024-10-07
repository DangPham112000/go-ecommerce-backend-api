# MVC architecture

Users -> Controller -> Service -> Repo -> Models -> DBs

# Testing

```cmd
go test -v
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
open coverage.html
```