package controllers

import "fmt"

// using an in memory store until I can get a DB or filesystem storage fired up
type PlayerStore interface {
	CreatePlayer(name string, score int) error
	GetPlayerScore(name string) (int, bool)
	UpdatePlayerScore(name string, newScore int) error
}

type PlayerServer struct {
	Store PlayerStore
}

type errUserExists struct{}

func (e *errUserExists) Error() string {
	return "user already exists"
}

type InMemoryPlayerStore struct {
	Scores map[string]int
}

func (store *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	result, exists := store.Scores[name]
	return result, exists
}

func (store *InMemoryPlayerStore) CreatePlayer(name string, score int) error {
	_, ok := store.Scores[name]
	if ok {
		return &errUserExists{}
	}
	store.Scores[name] = score
	return nil
}

func (store *InMemoryPlayerStore) UpdatePlayerScore(name string, newScore int) error {
	_, exists := store.Scores[name]

	if !exists {
		// user does not exist
		return &errUserExists{}
	}
	fmt.Println(name)
	store.Scores[name] = newScore
	return nil
}

func (store *InMemoryPlayerStore) DeletePlayer(name string, score int) error {
	_, exists := store.Scores[name]

	if exists {
		return &errUserExists{}
	}
	store.Scores[name] = score
	return nil
}

type AddPlayerRequest struct {
	Name  string
	Score int
}

type UpdatePlayerScoreRequest struct {
	Score int
}
