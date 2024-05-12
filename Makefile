include .env
url="postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable"
migration:
	atlas schema apply --url ${url} --dev-url "docker://postgres" --to "file://db/schema.sql" 

migrate:
	atlas migrate diff initial_schema --dir "file://migration" --to "file://db/schema.sql" --dev-url "docker://postgres"

generate:
	sqlc generate
	make migrate

run:
	sqlc generate
	make migrate
	air
	
