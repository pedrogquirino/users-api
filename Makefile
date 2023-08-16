run-api:
	@go run cmd/main.go

tidy: ## Download go dependencies
	@go mod tidy

go-test:
	@go test -p=1 -failfast ./...

setup-tests:
	@go install github.com/vektra/mockery/v2@2.20.0

docker-up:
	@docker-compose up -d --force-recreate

docker-down:
	@docker-compose down

documentation:
	@swag init -g cmd/main.go --outputTypes yaml