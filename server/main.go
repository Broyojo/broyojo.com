package main

import (
	"log"
	"net/http"
	"os"

	"github.com/caddyserver/certmagic"
)

func main() {
	fs := http.FileServer(http.Dir("../public"))
	if len(os.Args) >= 2 {
		if os.Args[1] == "online" {
			log.Println("Starting online server")
			log.Fatal(certmagic.HTTPS([]string{"broyojo.com", "www.broyojo.com"}, fs))
		}
	}
	log.Println("Starting local server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", fs))
}
