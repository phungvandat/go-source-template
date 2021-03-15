init-env: 
	@cat .env.example > .env

db-up:
	@docker-compose -f docker-compose-local.yml up -d

db-down:
	@docker-compose -f docker-compose-local.yml down

init-dev: init-env db-up

dev: migrate-db-up
	@go mod tidy
	@go run cmd/server/*.go

migrate-db-up:
	@cd cmd/migrator;\
	go run *.go
