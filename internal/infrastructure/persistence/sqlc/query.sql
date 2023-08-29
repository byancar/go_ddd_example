-- name: GetUserByID
SELECT * FROM user WHERE id = $1;

-- name: GetAllUsers
SELECT * FROM user;

-- name: InsertUser
INSERT INTO user (username, email) VALUES ($1, $2) RETURNING id;

-- name: UpdateUser
UPDATE user SET username = $2, email = $3 WHERE id = $1;

-- name: DeleteUser
DELETE FROM user WHERE id = $1;

-- name: GetAddressesByUserID
SELECT * FROM address WHERE user_id = $1;

-- name: InsertAddress
INSERT INTO address (user_id, street, city, state, country, postal_code)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;

-- name: UpdateAddress
UPDATE address
SET street = $2, city = $3, state = $4, country = $5, postal_code = $6
WHERE id = $1;

-- name: DeleteAddress
DELETE FROM address WHERE id = $1;