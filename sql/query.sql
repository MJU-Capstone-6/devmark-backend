-- name: FindUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: Create :one
INSERT INTO users (name) VALUES ($1) RETURNING *;
