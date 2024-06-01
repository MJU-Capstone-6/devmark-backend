-- +goose Up
ALTER TABLE "refresh_token" 
ALTER COLUMN "token" TYPE varchar(200);
-- +goose Down
ALTER TABLE "refresh_token" 
ALTER COLUMN "token" TYPE varchar;

