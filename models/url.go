package models

// URL represents a mapping between a shortened URL and its original URL.
type URL struct {
    ShortURL    string `json:"short_url"`   // The shortened version of the URL
    OriginalURL string `json:"original_url"` // The original full-length URL
}
