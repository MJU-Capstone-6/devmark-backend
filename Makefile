include .env
url="postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable"
migration:
	atlas schema apply --url ${url} --dev-url "docker://postgres" --to "file://db/schema.sql"

run:
	go run cmd/main.go
	
