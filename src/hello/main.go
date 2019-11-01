package main

import (
	"log"
	"net/http"
)

func main() {
        <broken code>
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
