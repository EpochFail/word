package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net"
	"net/http"
)

func UpVoteMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	word := r.URL.Query().Get(":word")
	updateRating(word, 1, r.RemoteAddr)

	log.Print("Up vote for: ", word)
}

func DownVoteMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	word := r.URL.Query().Get(":word")
	updateRating(word, -1, r.RemoteAddr)

	log.Print("Down vote for: ", word)
}

func updateRating(word string, vote int, address string) {
	host, _, err := net.SplitHostPort(address)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", ConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into votes (word, vote, ipaddress) values ($1, $2, $3)", word, vote, host)
	if err != nil {
		panic(err)
	}
}
