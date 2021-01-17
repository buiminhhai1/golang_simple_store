-- name: CreateBook :one
INSERT INTO books (
  name,
  price,
  quantity,
  book_category_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetBookById :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: GetBooksByCategoryId :many
SELECT * FROM books
WHERE book_category_id = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateBook :one
UPDATE books
SET name = $2, price = $3, quantity = $4, book_category_id = $5
WHERE id = $1
RETURNING *;

-- name: UpdateQuantityBook :one
UPDATE books
SET quantity = $2
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books WHERE id = $1;