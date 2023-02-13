package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func handleRequests() {
	mux := chi.NewRouter()
	mux.Get("/", allVideos)
	mux.Get("/getVideo/{id}", getVideo)
	mux.Post("/addVideo", addVideo)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
