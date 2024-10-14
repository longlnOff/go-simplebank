postgres:
	docker run --name postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -p 5432:5432 -d postgres:16-alpine

start_docker:
	docker start postgres

stop_docker:
	docker stop postgres

create_db:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

drop_db:
	docker exec -it postgres dropdb simple_bank

up_migrate:
	migrate -path sql/schema -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

down_migrate:
	migrate -path sql/schema -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test_all:
	go test -v -cover ./...

server:
	go run cmd/server/main.go


.phony: server test_all sqlc start_docker stop_docker postgres up_migrate down_migrate