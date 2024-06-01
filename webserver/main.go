package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/caddyserver/certmagic"
)

func main() {
	prod := flag.Bool("prod", false, "flag for production mode")
	flag.Parse()

	fs := http.FileServer(http.Dir("../site"))

	if *prod {
		log.Println("Starting HTTPS server")
		certmagic.DefaultACME.Agreed = true
		certmagic.DefaultACME.Email = "dha@xoba.com"
		log.Fatal(certmagic.HTTPS([]string{
			"www.broyojo.com",
			"broyojo.com",
			"www.davidhandrews.com",
			"davidhandrews.com",
		}, fs))
	} else {
		log.Println("Starting HTTP server on port 8080")
		log.Fatal(http.ListenAndServe(":8080", fs))
	}
}
