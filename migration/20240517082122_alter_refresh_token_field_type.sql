-- +goose Up
ALTER TABLE "refresh_token" 
ALTER COLUMN "token" TYPE text;

-- +goose Down
ALTER TABLE "refresh_token" 
ALTER COLUMN "token" TYPE varchar(300);


