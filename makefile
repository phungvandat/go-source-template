init-env: 
	@cat .env.example > .env
	@cat cmd/migrator/.env.migrator.example > cmd/migrator/.env.migrator

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

gen-docs: 
	swagger generate spec -m -o ./docs/api.yaml
	cp /dev/null "./docs/swagger.yaml"
	awk '{print}' "./docs/model.yaml" "./docs/api.yaml" >> "./docs/swagger.yaml"

docs-ui: 
	swagger serve ./docs/swagger.yaml