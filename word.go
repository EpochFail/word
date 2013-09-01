package main

import (
	"fmt"
	"github.com/bmizerany/pat"
	_ "github.com/lib/pq"
  "github.com/daaku/go.httpgzip"
	"database/sql"
	"net/http"
	"os"
  "log"
)

func main() {
	m := pat.New()
	m.Get("/api/word", httpgzip.NewHandler(http.HandlerFunc(wordMeBro)))
  m.Get("/api/vote/:word/up", httpgzip.NewHandler(http.HandlerFunc(upVoteMe)))  
  m.Get("/api/vote/:word/down", httpgzip.NewHandler(http.HandlerFunc(downVoteMe)))  

	http.Handle("/", m)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func wordMeBro(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

  db, err := sql.Open("postgres", "user=jsname password=test dbname=jsname_dev sslmode=disable")
  if err != nil {
    panic(err)
  }
		
  var word string
  var rating string
		
  db.QueryRow("SELECT word, rating FROM words OFFSET random()*(select max(word_id) from words) LIMIT 1").Scan(&word, &rating)

	fmt.Fprintf(w, "{\"word\":\"%v\", \"rating\":\"%v\"}", word, rating)
}

func upVoteMe(w http.ResponseWriter, r *http.Request) {
  word := r.URL.Query().Get(":word")
	updateRating(word, 1)

	log.Print("Up vote for: ", word)
}

func downVoteMe(w http.ResponseWriter, r *http.Request) {
  word := r.URL.Query().Get(":word")
	updateRating(word, -1)
	
	log.Print("Down vote for: ", word)
}

func updateRating(word string, vote int) {
  db, err := sql.Open("postgres", "user=jsname password=test dbname=jsname_dev sslmode=disable")
  if err != nil {
    panic(err)
  }
	
  queryString := fmt.Sprintf("insert into votes (word, vote) values ('%s', %d)", word, vote)
  db.Exec(queryString)
}
