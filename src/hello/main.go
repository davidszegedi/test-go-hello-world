package main

import (
	"log"
	"net/http"
)

func main() {
        <maybe another round broken code>
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
