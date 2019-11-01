package main

import (
	"log"
	"net/http"
)

func main() {
        <another round broken code>
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
