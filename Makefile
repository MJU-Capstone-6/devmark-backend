include .env
url="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"
migration_dir="./migration"

status:
	GOOSE_MIGRATION_DIR=${migration_dir} goose postgres ${url} status

up:
	GOOSE_MIGRATION_DIR=${migration_dir} goose postgres ${url} up

down:
	GOOSE_MIGRATION_DIR=${migration_dir} goose postgres ${url} down 


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
	
