package handlers

import (
    "encoding/json"
    "log"
    "math/rand"
    "net/http"
    "net/url"
    "strings"

    "github.com/gorilla/mux"
    "github.com/VGRobert/go-url-shortener/storage"
)

type ShortenRequest struct {
    URL string `json:"url"`
}

type ShortenResponse struct {
    ShortURL string `json:"short_url"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to the URL Shortener Service!"))
}

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
    var req ShortenRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Check if the URL has a scheme (http:// or https://). If not, add https:// by default.
    if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
        req.URL = "https://" + req.URL
    }

    // Validate the URL
    parsedURL, err := url.ParseRequestURI(req.URL)
    if err != nil {
        http.Error(w, "Invalid URL format", http.StatusBadRequest)
        return
    }

    log.Println("Shortening URL:", parsedURL.String())

    shortURL := generateShortURL()
    err = storage.Save(shortURL, parsedURL.String())
    if err != nil {
        http.Error(w, "Failed to save URL", http.StatusInternalServerError)
        return
    }

    res := ShortenResponse{ShortURL: shortURL}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    shortURL := vars["shortURL"]

    originalURL, found := storage.Load(shortURL)
    if !found {
        http.NotFound(w, r)
        return
    }

    http.Redirect(w, r, originalURL, http.StatusFound)
}

func generateShortURL() string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    const length = 6
    shortURL := make([]byte, length)
    for i := range shortURL {
        shortURL[i] = charset[rand.Intn(len(charset))]
    }
    return string(shortURL)
}
