-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
  to_account_id,
  amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransfers :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: GetTransfersByAccountID :one
SELECT * FROM transfers
WHERE to_account_id = $1 
OR from_account_id = $2;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListTransfersByAccountID :many
SELECT * FROM transfers
ORDER BY from_account_id
LIMIT $1
OFFSET $2;

-- name: DeleteTransfers :exec
DELETE FROM transfers
WHERE id = $1;

-- name: UpdateTransfers :exec
UPDATE transfers 
SET amount = $2
WHERE id = $1;