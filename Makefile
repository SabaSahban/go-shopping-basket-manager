#@IgnoreInspection BashAddShebang

export POSTGRES_ADDRESS=localhost:54320
export POSTGRES_DATABASE=basket
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=secret
export POSTGRES_DSN="postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_ADDRESS)/$(POSTGRES_DATABASE)?sslmode=disable"

migrate-create:
	migrate create -ext sql -dir ./migrations $(NAME)

migrate-up:
	migrate -verbose  -path ./migrations -database $(POSTGRES_DSN) up

migrate-down:
	 migrate -path ./migrations -database $(POSTGRES_DSN) down

migrate-reset:
	 migrate -path ./migrations -database $(POSTGRES_DSN) drop

run-server:
	go run . server

up:
	docker-compose up