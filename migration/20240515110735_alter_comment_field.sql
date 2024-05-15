-- +goose Up
ALTER TABLE "comment" 
ALTER COLUMN "context" TYPE varchar(300);

-- +goose Down
ALTER TABLE "comment" 
ALTER COLUMN "context" TYPE varchar;


