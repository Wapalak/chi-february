.PHONY: postgres adminer migrate

postgres:
	docker run --rm -ti -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -ti -p 8080:8080 adminer

migrate:
	migrate -source file://migrations -database postgres://postgres:12345@localhost/postgres?sslmode=disable up

migrate-down:
	migrate -source file://migrations -database postgres://postgres:12345@localhost/postgres?sslmode=disable down

server:
	go run .\cmd\chi_test_second\main.go
