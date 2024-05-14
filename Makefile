include .env
url="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

migration:
	atlas schema apply --url ${url} --dev-url "docker://postgres" --to "file://db/schema.sql" 

.PHONY: migration

generate:
	sqlc generate
	make migrate

run:
	sqlc generate
	make migrate
	air

swag:
	swag fmt
	swag init -d internal/auth,internal/category,internal/inviteCode,internal/workspace,internal/user,internal/bookmark -g ../../main.go --parseDependency
	
