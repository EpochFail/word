package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func UpVoteMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	word := r.URL.Query().Get(":word")
	updateRating(word, 1)

	log.Print("Up vote for: ", word)
}

func DownVoteMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	word := r.URL.Query().Get(":word")
	updateRating(word, -1)

	log.Print("Down vote for: ", word)
}

func updateRating(word string, vote int) {
	db, err := sql.Open("postgres", ConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into votes (word, vote) values ($1, $2)", word, vote)
	if err != nil {
		panic(err)
	}
}
