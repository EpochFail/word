package main

import (
	"github.com/alt234/word/util"
	"github.com/bmizerany/pat"
	"github.com/daaku/go.httpgzip"
	"net/http"
	"os"
)

func main() {
	m := pat.New()
	m.Get("/api/word", httpgzip.NewHandler(http.HandlerFunc(util.WordMeBro)))
	m.Get("/api/vote/:word/up", httpgzip.NewHandler(http.HandlerFunc(util.UpVoteMe)))
	m.Get("/api/vote/:word/down", httpgzip.NewHandler(http.HandlerFunc(util.DownVoteMe)))
	m.Get("/api/history", httpgzip.NewHandler(http.HandlerFunc(util.HistoryMe)))
	m.Get("/api/top10", httpgzip.NewHandler(http.HandlerFunc(util.Top10Me)))
	m.Get("/api/bottom10", httpgzip.NewHandler(http.HandlerFunc(util.Bottom10Me)))

	http.Handle("/", m)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
