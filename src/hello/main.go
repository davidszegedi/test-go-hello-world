package main

import (
	"log"
	"net/http"
)

func main() {
        <this time hopefully broken code>
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
