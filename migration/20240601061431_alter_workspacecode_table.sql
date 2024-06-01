-- +goose Up

ALTER TABLE workspace_code
ADD COLUMN user_id bigint;

ALTER TABLE workspace_code
ADD CONSTRAINT fk_user_id
FOREIGN KEY (user_id) REFERENCES "user"(id); 
-- +goose Down
ALTER TABLE workspace_code
DROP FOREIGN KEY fk_user_id;

ALTER TABLE workspace_code
DROP COLUMN user_id;
