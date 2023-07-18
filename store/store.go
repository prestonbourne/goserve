package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
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

func NewPostgresStore(ctx context.Context, postgresURL string) (*PostgresStore, error) {

	connConfig, err := pgx.ParseConfig(postgresURL)

	db, err := sql.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		return nil, fmt.Errorf("Failure]: Could not connect to Postgres %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("[Failure]: Could not Ping Postgres on Connection: %w", err)
	}

	fmt.Println("[Success]: Connected to Postgres")

	// create store
	store := &PostgresStore{
		db: db,
	}
	// init tables
	fmt.Println("start")
	store.createUsersTable()
	store.createTodosTable()
	fmt.Println("stop")
	return store, nil
}

// error handling, should i use context + should it be a pointer
func (s *PostgresStore) createUsersTable() error {
	const query string = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		user_name VARCHAR(14),
		created_at TIMESTAMP,
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
		last_edited TIMESTAMP
		created_at TIMESTAMP,
	);`
	if _, err := s.db.Exec(query); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil

}

func (s *PostgresStore) AddUser() error {
	return nil
}
