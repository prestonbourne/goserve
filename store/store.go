package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/prestonbourne/goserve/utils"
)

// Johan's original schema
// const schema = `
// CREATE TABLE players (
//
//	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//	name TEXT NOT NULL,
//	score INTEGER NOT NULL DEFAULT 0 CHECK (score >= 0)
//
// );
// `
// might be a good option to learn singletons in GO to ensure one DB connection
type PostgresStore struct {
	db *sql.DB
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	UserName  string    `json:"userName"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewPostgresStore(ctx context.Context, postgresURL string) (*PostgresStore, error) {

	connConfig, err := pgx.ParseConfig(postgresURL)

	db, err := sql.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		return nil, fmt.Errorf("Failure]: Could not connect to Postgres %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("[Failure]: Could not Ping Postgres on Connection: %w", err)
	}

	utils.Success("Connected to Postgres")

	// create store
	store := &PostgresStore{
		db: db,
	}
	// init tables

	store.createUsersTable()
	store.createTodosTable()

	return store, nil
}

// error handling, should i use context + should it be a pointer
func (s *PostgresStore) createUsersTable() error {
	// no trailing commas in SQL query!
	const query string = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		user_name VARCHAR(14),
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
	if _, err := s.db.Exec(query); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (s *PostgresStore) createTodosTable() error {
	const query string = `CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		is_complete BOOLEAN,
		content TEXT,
		last_edited TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);`
	if _, err := s.db.Exec(query); err != nil {
		fmt.Println("%w", err)
		return fmt.Errorf("%w", err)
	}

	return nil

}

/*
having to pass all the params is...pretty ugly,
i can't import the user type because circular dependencies
how can i make this project structure cleaner 🤔
*/

func (s *PostgresStore) AddUser(firstName string, lastName string, userName string, createdAt time.Time) error {

	const query string = `INSERT INTO users
(first_name, last_name, user_name, created_at)
VALUES ($1, $2, $3, $4)`

	resp, err := s.db.Exec(query, firstName, lastName, userName, createdAt)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgresStore) GetUsers() ([]*User, error) {
	rows, err := s.db.Query("SELECT * FROM users;")

	if err != nil {
		return nil, err
	}

	users := []*User{}
	// can this loop be generalized?
	for rows.Next() {
		user, err := scanIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
func (s *PostgresStore) GetUserByID(id int) (*User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1;", id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUser(rows)

	}
	return nil, fmt.Errorf("Could not find User with id of %v", id)
}

func (s *PostgresStore) GetUserByUsername(username string) (*User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE username = $1;", username)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUser(rows)

	}
	return nil, fmt.Errorf("Could not find User: %v", username)
}

/*
Note to future self: explore Soft vs Hard deleting if
you really get into this and build a UI etc
*/
func (s *PostgresStore) DeleteUser(id int) error {
	_, err := s.db.Query("DELETE FROM users WHERE id = $1;", id)

	if err != nil {
		return fmt.Errorf("Could not find User with id of %v", id)
	}

	return nil
}

func scanIntoUser(rows *sql.Rows) (*User, error) {

	// can this loop be generalized?
	user := &User{}

	err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.UserName, &user.CreatedAt)
	return user, err
}
