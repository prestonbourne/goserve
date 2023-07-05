package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// using an in memory store until I can get a DB or filesystem storage fired up
type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
}

type PlayerServer struct {
	store PlayerStore
}

type InMemoryPlayerStore struct {
	scores map[string]int
}

func (store *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	result, exists := store.scores[name]
	return result, exists
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(http.StatusAccepted)
	case http.MethodGet:
		player := strings.TrimPrefix(r.URL.Path, "/players/")
		score, exists := p.store.GetPlayerScore(player)

		if exists != false {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		fmt.Fprint(w, score)
	}

}
