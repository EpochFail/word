package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net"
	"net/http"
	"strings"
)

func UpVoteMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	word := r.URL.Query().Get(":word")
	updateRating(word, 1, r)

	log.Print("Up vote for: ", word)
}

func DownVoteMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	word := r.URL.Query().Get(":word")
	updateRating(word, -1, r)

	log.Print("Down vote for: ", word)
}

func updateRating(word string, vote int, r *http.Request) {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		panic(err)
	}

	realIp := r.Header.Get("X-Real-Ip")
	forwardedFor := r.Header.Get("X-Forwarded-For")
	
  if realIp != "" {
		host = realIp
	}

	if forwardedFor != "" {
    addresses := strings.Split(forwardedFor, ",")
		for _, ele := range addresses {
			if ele != "127.0.0.1" {
				host = strings.TrimSpace(ele)
        break
			}
		}
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
