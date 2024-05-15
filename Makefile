include .env
url="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"
migration_dir="./migration"

AUTH_PATH="internal/auth"
CATEGORY_PATH="internal/category"
INVITECODE_PATH="internal/inviteCode"
WORKSPACE_PATH="internal/workspace"
USER_PATH="internal/user"
BOOKMARK_PATH="internal/bookmark"
REFRESH_TOKEN_PATH="internal/refreshToken"
COMMENT_PATH="internal/comment"

status:
	GOOSE_MIGRATION_DIR=${migration_dir} goose postgres ${url} status

up:
	GOOSE_MIGRATION_DIR=${migration_dir} goose postgres ${url} up

down:
	GOOSE_MIGRATION_DIR=${migration_dir} goose postgres ${url} down 


generate:
	sqlc generate
	make up

run:
	make generate
	air

swag:
	swag fmt
	swag init -d ${AUTH_PATH},${CATEGORY_PATH},${INVITECODE_PATH},${WORKSPACE_PATH},${USER_PATH},${BOOKMARK_PATH},${REFRESH_TOKEN_PATH},${COMMENT_PATH} -g ../../main.go --parseDependency
	
