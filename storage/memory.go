package storage

import "sync"

type MemoryStore struct {
    urls map[string]string
    sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
    return &MemoryStore{
        urls: make(map[string]string),
    }
}

func (s *MemoryStore) Save(shortURL, originalURL string) {
    s.Lock()
    defer s.Unlock()
    s.urls[shortURL] = originalURL
}

func (s *MemoryStore) Load(shortURL string) (string, bool) {
    s.RLock()
    defer s.RUnlock()
    originalURL, found := s.urls[shortURL]
    return originalURL, found
}
