package main

import (
	"log"
	"net/http"
)

func main() {
        <maybe1 another round broken code>
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
