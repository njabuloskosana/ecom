build:
	@go build -o bin/ecom go/api/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/ecom

migration:
	@migrate create -ext sql -dir go/db/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run go/db/migrate/main.go up

migrate-down:
	@go run go/db/migrate/main.go down