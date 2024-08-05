-- name: CreateUser :one
INSERT INTO users (
  id,
  username,
  email,
  email_verified,
  state
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: CreateUserCredential :one
INSERT INTO user_credentials (
  user_id,
  credential,
  salt
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: CreateUserSocial :one
INSERT INTO user_social_tokens (
  user_id,
  provider_type,
  access_token,
  refresh_token,
  expires
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUserById :one
Select * from users where id = $1;

-- name: GetAllUsers :many
Select * from users;

-- name: GetUserByUsername :one
Select * from users where username = $1;

-- name: GetUserByEmail :one
Select * from users where email = $1;

-- name: GetUserByUsernameAndVerifyPassword :one
Select * from users u where u.username = $1 and u.email_verified = $2;

-- name: GetUserCredentialByUserId :one
Select * from user_credentials u where u.user_id = $1;