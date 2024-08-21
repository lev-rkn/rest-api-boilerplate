
cover:
	go test -short -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

swag_init:
	swag init -g internal/controllers/router.go
	
run-local:
	go run ./cmd/main.go --config=local

migrations-up:
	goose -dir ./migrations postgres \
	postgres://postgres:postgres@localhost:5432/rest-api-service?sslmode=disable \
	up
