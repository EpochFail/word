package main

import (
	"bufio"
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/daaku/go.httpgzip"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	file, err := os.Open("result.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	rand.Seed(time.Now().UnixNano())
	max := len(lines)

	wordMeBro := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		word := lines[rand.Intn(max)]
		fmt.Fprintf(w, "{\"word\":\"%v\"}", word)
	}

	m := pat.New()
	m.Get("/api/word", httpgzip.NewHandler(http.HandlerFunc(wordMeBro)))
	http.Handle("/", m)

	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
