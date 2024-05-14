include .env
url="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

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

swag:
	swag fmt
	swag init -d internal/auth,internal/category,internal/inviteCode,internal/workspace,internal/user -g ../../main.go --parseDependency
	
