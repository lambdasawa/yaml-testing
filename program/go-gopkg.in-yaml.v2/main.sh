[ -z "$BUILD" ] || go mod tidy
go run main.go $1
