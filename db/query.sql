-- name: CreateBook :one
INSERT INTO book (
  title, author
) VALUES (
  ?, ?
)
RETURNING *;

-- name: GetAllBooks :many
SELECT *
FROM book;
