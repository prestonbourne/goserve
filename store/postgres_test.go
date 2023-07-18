package store

import (
	"context"
	"testing"
)

func TestNewPostgresStore(t *testing.T) {
	_, err := NewPostgresStore(context.Background(), "postgresql://postgres:password@localhost:5432/postgres")
	if err != nil {
		t.Fatal("Failed to create store:", err)
	}
}
