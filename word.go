package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/bmizerany/pat"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

var ConnectionString = "user=word password=krampus dbname=word_dev sslmode=disable"

func main() {
	m := pat.New()
	m.Get("/api/word", http.HandlerFunc(WordMeBro))
	m.Get("/api/vote/:word/up", http.HandlerFunc(UpVoteMe))
	m.Get("/api/vote/:word/down", http.HandlerFunc(DownVoteMe))
	m.Get("/api/history", http.HandlerFunc(HistoryMe))
	m.Get("/api/top10", http.HandlerFunc(Top10Me))
	m.Get("/api/bottom10", http.HandlerFunc(Bottom10Me))
	m.Get("/api/random10", http.HandlerFunc(Random10Me))

	http.Handle("/", m)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func WordMeBro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("postgres", ConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var word_id int
	var word string
	var rating int

	db.QueryRow("SELECT word_id, word, rating FROM words OFFSET random()*(select max(word_id) from words) LIMIT 1").Scan(&word_id, &word, &rating)

	_, err = db.Exec("insert into word_history (word_id) values ($1)", word_id)
	if err != nil {
		panic(err)
	}

	m := wordList{word, rating}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(b[:]))
}

type wordList struct {
	Word   string
	Rating int
}

func Random10Me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("postgres", ConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select word, rating from words order by random() limit 10")
	if err != nil {
		panic(err)
	}

	var random10 []wordList
	for rows.Next() {
		var _word string
		var _rating int

		rows.Scan(&_word, &_rating)

		random10 = append(random10, wordList{_word, _rating})
	}

	t, err := json.Marshal(random10)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(t[:]))
}
