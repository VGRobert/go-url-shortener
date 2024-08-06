package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "github.com/VGRobert/go-url-shortener/handlers"
    "github.com/VGRobert/go-url-shortener/storage"
)

func main() {
    storage.InitDB()

    r := mux.NewRouter()

    // Routes
    r.HandleFunc("/shorten", handlers.ShortenURLHandler).Methods("POST")
    r.HandleFunc("/{shortURL}", handlers.RedirectHandler).Methods("GET")

    // Enable CORS
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"}, // Adjust to your frontend URL
        AllowedMethods:   []string{"GET", "POST"},
        AllowedHeaders:   []string{"Content-Type"},
        AllowCredentials: true,
    })

    // Start the server with CORS enabled
    log.Println("Starting the server on :8080...")
    log.Fatal(http.ListenAndServe(":8080", c.Handler(r)))
}
