-- name: CreateOrder :one
INSERT INTO orders (
  user_id,
  status
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetOrderById :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: GetOrdersByUserId :many
SELECT * FROM orders
WHERE user_id = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListOrders :many
SELECT * FROM orders
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateOrderStatus :one
UPDATE orders
SET status = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOrderById :exec
DELETE FROM orders WHERE id = $1;