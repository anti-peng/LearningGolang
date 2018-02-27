package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("/Users/fry/GreatLD/TestNeo4j/pipe.json"))
	http.Handle("/json", fs)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
