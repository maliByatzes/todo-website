package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/maliByatzes/todo-website/util"
	"github.com/stretchr/testify/require"
)

func createRandomTodo(t *testing.T) Todo {
	user := createRandomUser(t)

	arg := CreateTodoParams{
		Username:    user.Username,
		TodoName:    util.RandomString(10),
		Description: util.RandomString(40),
		IsCompleted: false,
	}

	todo, err := testQueries.CreateTodo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.Equal(t, arg.Username, todo.Username)
	require.Equal(t, arg.TodoName, todo.TodoName)
	require.Equal(t, arg.Description, todo.Description)
	require.Equal(t, arg.IsCompleted, todo.IsCompleted)

	require.True(t, todo.UpdatedAt.IsZero())
	require.NotZero(t, todo.CreatedAt)

	return todo
}

func TestCreateTodo(t *testing.T) {
	createRandomTodo(t)
}

func TestGetTodo(t *testing.T) {
	todo1 := createRandomTodo(t)
	require.NotEmpty(t, todo1)

	todo2, err := testQueries.GetTodo(context.Background(), todo1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, todo2)

	require.Equal(t, todo1.ID, todo2.ID)
	require.Equal(t, todo1.Username, todo2.Username)
	require.Equal(t, todo1.TodoName, todo2.TodoName)
	require.Equal(t, todo1.Description, todo2.Description)
	require.Equal(t, todo1.IsCompleted, todo2.IsCompleted)
	require.WithinDuration(t, todo1.UpdatedAt, todo2.UpdatedAt, time.Second)
	require.WithinDuration(t, todo1.CreatedAt, todo2.CreatedAt, time.Second)
}

func TestUpdateTodo(t *testing.T) {
	todo1 := createRandomTodo(t)
	require.NotEmpty(t, todo1)

	arg := UpdateTodoParams{
		ID:          todo1.ID,
		TodoName:    util.RandomString(10),
		Description: util.RandomString(40),
		IsCompleted: true,
	}

	todo2, err := testQueries.UpdateTodo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo2)

	require.Equal(t, todo1.ID, todo2.ID)
	require.Equal(t, todo1.Username, todo2.Username)
	require.Equal(t, arg.TodoName, todo2.TodoName)
	require.Equal(t, arg.Description, todo2.Description)
	require.Equal(t, arg.IsCompleted, todo2.IsCompleted)
	require.WithinDuration(t, todo1.UpdatedAt, todo2.UpdatedAt, time.Second)
	require.WithinDuration(t, todo1.CreatedAt, todo2.CreatedAt, time.Second)
}

func TestDeleteTodo(t *testing.T) {
	todo1 := createRandomTodo(t)

	err := testQueries.DeleteTodo(context.Background(), todo1.ID)
	require.NoError(t, err)

	todo2, err := testQueries.GetTodo(context.Background(), todo1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, todo2)
}

func TestListTodos(t *testing.T) {
	var lastTodo Todo
	for i := 0; i < 15; i++ {
		lastTodo = createRandomTodo(t)
	}

	arg := ListTodosParams{
		Username: lastTodo.Username,
		Limit:    5,
		Offset:   0,
	}

	todos, err := testQueries.ListTodos(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todos)

	for _, todo := range todos {
		require.NotEmpty(t, todo)
		require.Equal(t, lastTodo.Username, todo.Username)
	}

}
