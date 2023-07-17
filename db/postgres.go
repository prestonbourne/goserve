package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

const schema = `
CREATE TABLE players (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name TEXT NOT NULL,
	score INTEGER NOT NULL DEFAULT 0 CHECK (score >= 0)
);
`

func Init(ctx context.Context, dbURL string) (*sql.DB, error) {
	connConfig, err := pgx.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DB URL: %w", err)
	}

	db, err := sql.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping DB on connect: %w", err)
	}
	// if _, err := db.ExecContext(ctx, schema); err != nil {
	// 	return nil, fmt.Errorf("failed to apply schema: %w", err)
	// }

	fmt.Println("[success]: Connected to postgres!")
	return db, nil
}
