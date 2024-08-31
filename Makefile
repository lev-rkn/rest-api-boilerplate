cover:
	go test -short -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

swag-init:
	swag init -g internal/controller/router.go --output api/swagger
	
migrations-up:
	goose -dir ./migrations postgres \
	postgres://postgres:postgres@localhost:5432/rest-api-service?sslmode=disable \
	up

run-local:
	go run ./cmd/main.go --config=local

