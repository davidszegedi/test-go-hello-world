package main

import (
	"log"
	"net/http"
)

func main() {
        <even more broken code>
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
