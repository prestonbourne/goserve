package todos

import (
	"database/sql"
	"fmt"
	"time"
)

type TodoModel struct {
	Complete  bool      `json:"IsCompleted"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type TodoServer struct {
	DB *sql.DB
}

func (s *TodoServer) AddTodo(todo TodoModel) (TodoModel, error) {
	var out string
	queryString := fmt.Sprintf("INSERT INTO Todos (Content, CreatedAt, Complete) VALUES ('%v', '%v', '%v);", todo.Complete, todo.CreatedAt, todo.Complete)

	if res, err := s.DB.Exec(queryString); err != nil {
		return todo, fmt.Errorf("failed to ping DB on connect: %w", err)

	}

}
