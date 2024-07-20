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