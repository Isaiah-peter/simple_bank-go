-- name: CreateEntries :one
INSERT INTO entries (
  account_id,
  amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntries :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: GetEntriesByAccountID :one
SELECT * FROM entries
WHERE account_id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListEntriesByAccountID :many
SELECT * FROM entries
ORDER BY account_id
LIMIT $1
OFFSET $2;

-- name: DeleteEntries :exec
DELETE FROM entries
WHERE id = $1;

-- name: UpdateEntries :exec
UPDATE entries 
SET amount = $2
WHERE id = $1;