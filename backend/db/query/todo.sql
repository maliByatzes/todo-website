-- name: CreateTodo :one
INSERT INTO todos (
  username,
  todo_name,
  description,
  is_completed
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
WHERE username = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateTodo :one
UPDATE todos
SET 
  todo_name = $2, 
  description = $3, 
  is_completed = $4
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;