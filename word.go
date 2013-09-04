package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/daaku/go.httpgzip"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	m := pat.New()
	m.Get("/api/word", httpgzip.NewHandler(http.HandlerFunc(wordMeBro)))
	m.Get("/api/vote/:word/up", httpgzip.NewHandler(http.HandlerFunc(upVoteMe)))
	m.Get("/api/vote/:word/down", httpgzip.NewHandler(http.HandlerFunc(downVoteMe)))
	m.Get("/api/stats", httpgzip.NewHandler(http.HandlerFunc(statsMe)))

	http.Handle("/", m)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

type Message struct {
	Word   string
	Rating int
}

func wordMeBro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("postgres", "user=jsname password=test dbname=jsname_dev sslmode=disable")
	if err != nil {
		panic(err)
	}

	var word_id int
	var word string
	var rating int

	db.QueryRow("SELECT word_id, word, rating FROM words OFFSET random()*(select max(word_id) from words) LIMIT 1").Scan(&word_id, &word, &rating)

	result, err := db.Exec("insert into word_history (word_id) values ($1)", word_id)
	if err != nil {
		panic(err)
	}

	log.Print(result)

	m := Message{word, rating}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(b[:]))
}

func upVoteMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	word := r.URL.Query().Get(":word")
	updateRating(word, 1)

	log.Print("Up vote for: ", word)
}

func downVoteMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	word := r.URL.Query().Get(":word")
	updateRating(word, -1)

	log.Print("Down vote for: ", word)
}

type Stats struct {
	History  []Message
	Top10    []Message
	Bottom10 []Message
}

func statsMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	history := getStatsSlice("select w.word, w.rating from words w inner join word_history wh on w.word_id=wh.word_id order by wh.created_at desc limit 10")
	top10 := getStatsSlice("select distinct word, rating from words order by rating desc limit 10")
	bottom10 := getStatsSlice("select distinct word, rating from words order by rating limit 10")

	s := Stats{history, top10, bottom10}
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(b[:]))
}

func getStatsSlice(queryString string) []Message {
	// TODO: Stop opening db so often. Probably should try and ORM too.
	db, err := sql.Open("postgres", "user=jsname password=test dbname=jsname_dev sslmode=disable")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(queryString)
	if err != nil {
		panic(err)
	}

	var stats []Message
	for rows.Next() {
		var _word string
		var _rating int

		rows.Scan(&_word, &_rating)

		stats = append(stats, Message{_word, _rating})
	}

	return stats
}

func updateRating(word string, vote int) {
	db, err := sql.Open("postgres", "user=jsname password=test dbname=jsname_dev sslmode=disable")
	if err != nil {
		panic(err)
	}

	log.Print(word)
	log.Print(vote)

	result, err := db.Exec("insert into votes (word, vote) values ($1, $2)", word, vote)
	if err != nil {
		panic(err)
	}

	log.Print(result)
}
