-- name: CreateCategory :one
INSERT INTO book_categories (name) VALUES (
  $1
)
RETURNING *;

-- name: GetCategory :one
SELECT * FROM book_categories
WHERE id = $1 LIMIT 1;

-- name: ListBookCategories :many
SELECT * FROM book_categories
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCategory :one
UPDATE book_categories SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM book_categories WHERE id = $1;