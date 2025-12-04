package cache

import (
	"sync"
)

type SearchCache struct {
	results map[string]string
	mu      sync.Mutex
}

func NewSearchCache() *SearchCache {
	return &SearchCache{
		results: make(map[string]string),
	}
}

func (s *SearchCache) Get(id string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	result, ok := s.results[id]
	return result, ok
}

func (s *SearchCache) Set(id string, result string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.results[id] = result
}
