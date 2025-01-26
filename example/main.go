package main

import (
	"log"
	"net/http"

	"github.com/petrolax/sse-server-example"
)

func main() {
	server := http.NewServeMux()
	service := sse.NewService()
	hs := sse.NewHandlers(service, sse.RetryDuration)

	server.HandleFunc("/listen", hs.ListeHandler)
	server.HandleFunc("/say", hs.SayHandler)

	log.Fatalln(http.ListenAndServe(":8080", server))
}
