package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

type message struct {
	Word   string
	Rating int
}

func HistoryMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	history := getStatsSlice("select w.word, w.rating from words w inner join word_history wh on w.word_id=wh.word_id order by wh.created_at desc limit 10")

	b, err := json.Marshal(history)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(b[:]))
}

func Top10Me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	top10 := getStatsSlice("select distinct word, rating from words order by rating desc limit 10")

	b, err := json.Marshal(top10)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(b[:]))
}

func Bottom10Me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	bottom10 := getStatsSlice("select distinct word, rating from words order by rating limit 10")

	b, err := json.Marshal(bottom10)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(b[:]))
}

func getStatsSlice(queryString string) []message {
	db, err := sql.Open("postgres", ConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(queryString)
	if err != nil {
		panic(err)
	}

	var stats []message
	for rows.Next() {
		var _word string
		var _rating int

		rows.Scan(&_word, &_rating)

		stats = append(stats, message{_word, _rating})
	}

	return stats
}
