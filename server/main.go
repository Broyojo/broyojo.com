package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("../public"))
	log.Fatal(http.ListenAndServe(":8080", fileServer))
}
