package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/prestonbourne/goserve/db"
)

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("received")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		// POST /players

		var jsonRequest AddPlayerRequest
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&jsonRequest)
		if err != nil {
			http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		if jsonRequest.Name == "" {
			http.Error(w, "invalid player name", http.StatusBadRequest)
			return
		}
		if jsonRequest.Score < 0 {
			http.Error(w, "invalid player score", http.StatusBadRequest)
			return
		}
		err = p.Store.CreatePlayer(jsonRequest.Name, jsonRequest.Score)
		p.DB.AddPlayer(db.Player(jsonRequest))
		if err != nil {
			var existErr errUserExists
			if errors.Is(err, &existErr) {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case http.MethodGet:
		// GET /players/<name/id>

		if r.URL.Path == "/players" || r.URL.Path == "/players/" {
			jsonData, err := json.Marshal(p.Store)
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Fprintf(w, string(jsonData))
		}

		player := strings.TrimPrefix(r.URL.Path, "/players/")

		if player != r.URL.Path {
			score, exists := p.Store.GetPlayerScore(player)
			fmt.Println(exists)
			if !exists {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "Player Does not exist")
				return
			}

			fmt.Fprintf(w, player+"has "+fmt.Sprint(score)+" points")
		}
		p.DB.GetPlayers()
	case http.MethodPatch, http.MethodPut:

		var jsonRequest UpdatePlayerScoreRequest

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&jsonRequest)

		_, exists := p.Store.GetPlayerScore(jsonRequest.Name)
		if !exists {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Player Does not exist")
			return
		}

		if err == nil {
			p.Store.UpdatePlayerScore(jsonRequest.Name, jsonRequest.Score)
			newScore, _ := p.Store.GetPlayerScore(jsonRequest.Name)
			// think of a better name since I'm using the same type for res response
			res := UpdatePlayerScoreRequest{Name: jsonRequest.Name, Score: newScore}
			json.NewEncoder(w).Encode(res)

		}

	}
}
