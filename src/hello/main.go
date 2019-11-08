package main

import (
	"log"
	"net/http"
)

func main() {
        < error code >
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
