package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/epochfail/word"
)

func main() {
	var db word.DB
	if err := db.Open(); err != nil {
		log.Fatal(err)
	}
	server := word.NewHTTPServer(&db)
	router, err := word.CreateRouter(server)
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Handle("/", router)

	u := "0.0.0.0:8001"
	log.Printf("Word server started at http://%s\n", u)
	err = http.ListenAndServe(u, nil)
	if err != nil {
		fmt.Println(err)
	}
}
