package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var connectionString = "user=jsname password=test dbname=jsname_dev sslmode=disable"

type message struct {
	Word   string
	Rating int
}

func WordMeBro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("postgres", connectionString)
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
	m := message{word, rating}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(b[:]))
}

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
	db, err := sql.Open("postgres", connectionString)
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

func updateRating(word string, vote int) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into votes (word, vote) values ($1, $2)", word, vote)
	if err != nil {
		panic(err)
	}
}
