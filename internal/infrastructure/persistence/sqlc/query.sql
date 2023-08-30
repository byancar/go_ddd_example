-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: InsertUser :exec
INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id;

-- name: UpdateUser :exec
UPDATE users SET username = $2, email = $3 WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: GetAddressesByUserID :one
SELECT * FROM addresses WHERE user_id = $1;

-- name: InsertAddress :exec
INSERT INTO addresses (user_id, street, city, state, country, postal_code)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;

-- name: UpdateAddress :exec
UPDATE addresses
SET street = $2, city = $3, state = $4, country = $5, postal_code = $6
WHERE id = $1;

-- name: DeleteAddress :exec
DELETE FROM addresses WHERE id = $1;