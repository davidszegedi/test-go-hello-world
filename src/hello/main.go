package main

import (
	"log"
	"net/http"
)

func main() {
        <kinda broken code>
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
