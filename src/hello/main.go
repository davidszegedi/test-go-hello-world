package main

import (
	"log"
	"net/http"
)

func main() {
        <should work for the test but kinda broken code>
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
