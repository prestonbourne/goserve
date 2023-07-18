package todos

import (
	"math/rand"

	"time"
)

type Todo struct {
	ID         int       `json:"id"`
	Complete   bool      `json:"isComplete"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
	LastEdited time.Time `json:"lastEdited"`
}

func NewTodo(content string) *Todo {
	//TODO: format data properly, look into SQL & POSTGRES data types
	//TODO: Find a good system for ID's, how can we expose the ID to the client, should we?
	currentTime := time.Now()

	return &Todo{ID: rand.Intn(10000), Complete: false, Content: content, CreatedAt: currentTime, LastEdited: currentTime}
}

type TodoModel interface {
	Add(todo Todo) (Todo, error)
	GetById(id string) (Todo, error)
	Delete(id string) (Todo, error)
	Update(id string) (Todo, error)
}

// func (s *TodoServer) AddTodo(todo Todo) (Todo, error) {

// 	queryString := fmt.Sprintf("INSERT INTO Todos (Content, CreatedAt, Complete) VALUES ('%v', '%v', '%v);", todo.Complete, todo.CreatedAt, todo.Complete)

// 	if _, err := s.DB.Exec(queryString); err != nil {
// 		return todo, fmt.Errorf("failed to ping DB on connect: %w", err)

// 	}
// 	return todo, nil
// }
