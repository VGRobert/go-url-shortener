package storage

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
    var err error
    db, err = sql.Open("sqlite3", "./urls.db")
    if err != nil {
        log.Fatal(err)
    }

    createTable := `
    CREATE TABLE IF NOT EXISTS url_mapping (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        short_url TEXT NOT NULL UNIQUE,
        original_url TEXT NOT NULL
    );`

    _, err = db.Exec(createTable)
    if err != nil {
        log.Fatal(err)
    }
}

func Save(shortURL, originalURL string) error {
    _, err := db.Exec("INSERT INTO url_mapping (short_url, original_url) VALUES (?, ?)", shortURL, originalURL)
    return err
}

func Load(shortURL string) (string, bool) {
    var originalURL string
    err := db.QueryRow("SELECT original_url FROM url_mapping WHERE short_url = ?", shortURL).Scan(&originalURL)
    if err != nil {
        if err == sql.ErrNoRows {
            return "", false
        }
        log.Println("Error loading URL:", err)
        return "", false
    }
    return originalURL, true
}
