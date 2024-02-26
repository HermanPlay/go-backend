up:
	docker compose up --build -d
down:
	docker compose down
test:
	docker compose up --build -d
	docker compose exec -T api go test -cover ./internal/api/http
	docker compose down
unit-test:
	go test -coverprofile=c.out ./...
open-cov: unit-test
	go tool cover -html="c.out"
test-cov: postgres_up unit-test postgres_down
postgres_up:
	@docker run -d \
		--name postgres-container \
		-e POSTGRES_DB=backend \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		-p 5432:5432 \
		postgres:latest
postgres_down:
	@docker stop postgres-container
	@docker rm postgres-container