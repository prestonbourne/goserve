package db

import (
	"database/sql"
	"fmt"
	"log"
)

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type PostgresStore struct {
	db *sql.DB
}

func (store *PostgresStore) GetPlayers() {
	var out string
	res, err := store.db.Query("SELECT * FROM players")
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	res.Scan(&out)
	fmt.Printf("%v\n", out)
}

func (store *PostgresStore) AddPlayer(player Player) {

	var out string
	queryString := fmt.Sprintf("INSERT INTO players (name, score) VALUES ('%v', '%v');", player.Name, player.Score)

	res, err := store.db.Query(queryString)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	res.Scan(&out)
	fmt.Printf("%v\n", out)

}
