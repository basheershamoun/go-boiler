-- name: ListUsers :many
SELECT * FROM user WHERE deleted_at IS NULL;

-- name: GetUserById :one
SELECT * FROM user WHERE id = ? AND deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT * FROM user WHERE email = ? AND deleted_at IS NULL;

-- name: CreateUser :execresult
INSERT INTO user (uuid, roles, password, email, username, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NOW(), NOW());

-- name: UpdateUser :execresult
UPDATE user SET roles = ?, password = ?, email = ?, username = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteUser :execresult
UPDATE user SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL;


