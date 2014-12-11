package main

import (
	"bufio"
	"database/sql"
	"os"
)

// Up is executed when this migration is applied
func Up_20131009224738(txn *sql.Tx) {
	file, err := os.Open("seeds/wordlist.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txn.Exec("insert into words (word, rating) values ($1,0)", scanner.Text())
	}
}

// Down is executed when this migration is rolled back
func Down_20131009224738(txn *sql.Tx) {
	txn.Exec("truncate table words")
}
