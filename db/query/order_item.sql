-- name: CreateOrderItem :one
INSERT INTO order_items (
  order_id,
  book_id,
  quantity,
  price_invoice
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetOrderItemByPrimaryKey :one
SELECT * FROM order_items
WHERE order_id = $1 AND book_id = $2
LIMIT 1;

-- name: GetOrderItemsByBookID :many
SELECT * FROM order_items
WHERE book_id = $3 
LIMIT $1
OFFSET $2;

-- name: GetOrderItemsByOrderID :many
SELECT * FROM order_items
WHERE order_id = $3
ORDER BY book_id
LIMIT $1
OFFSET $2;

-- name: UpdateOrderPriceInvoice :one
UPDATE order_items
SET price_invoice = $3
WHERE order_id = $1 AND book_id = $2
RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM order_items WHERE order_id = $1 AND book_id = $2;