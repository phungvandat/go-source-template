dev: migrate-db-up
	go mod tidy
	go run cmd/server/*.go

migrate-db-up:
	cd cmd/migrator;\
	go run *.go
