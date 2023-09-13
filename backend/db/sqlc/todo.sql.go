// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: todo.sql

package db

import (
	"context"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (
  username,
  todo_name,
  description,
  is_completed
) VALUES (
  $1, $2, $3, $4
) RETURNING id, username, todo_name, description, is_completed, updated_at, created_at
`

type CreateTodoParams struct {
	Username    string `json:"username"`
	TodoName    string `json:"todo_name"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo,
		arg.Username,
		arg.TodoName,
		arg.Description,
		arg.IsCompleted,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.TodoName,
		&i.Description,
		&i.IsCompleted,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getTodo = `-- name: GetTodo :one
SELECT id, username, todo_name, description, is_completed, updated_at, created_at FROM todos
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.TodoName,
		&i.Description,
		&i.IsCompleted,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, username, todo_name, description, is_completed, updated_at, created_at FROM todos
WHERE username = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListTodosParams struct {
	Username string `json:"username"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

func (q *Queries) ListTodos(ctx context.Context, arg ListTodosParams) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodos, arg.Username, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Todo{}
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.TodoName,
			&i.Description,
			&i.IsCompleted,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos
SET 
  todo_name = $2, 
  description = $3, 
  is_completed = $4
WHERE id = $1
RETURNING id, username, todo_name, description, is_completed, updated_at, created_at
`

type UpdateTodoParams struct {
	ID          int64  `json:"id"`
	TodoName    string `json:"todo_name"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodo,
		arg.ID,
		arg.TodoName,
		arg.Description,
		arg.IsCompleted,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.TodoName,
		&i.Description,
		&i.IsCompleted,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}