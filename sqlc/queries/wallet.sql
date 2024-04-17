-- name: CreateWallet :exec
INSERT INTO wallet (
  id,
  user_id,
  address,
  network, 
  asset
) VALUES (
    $1,$2,$3,$4,$5
);

-- name: CreateInternalWallet :exec
INSERT INTO internal_wallet(
  wallet_id, 
  derivation_version,
  index
)
VALUES ($1, $2,$3);

-- name: GetAllWallet :many
SELECT * FROM wallet 
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetWallet :one
SELECT * 
FROM wallet AS w
RIGHT JOIN internal_wallet AS i ON i.wallet_id = w.id
WHERE w.user_id = $1 AND w.network = $2 AND w.asset = $3;

