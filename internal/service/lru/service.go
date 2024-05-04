package lru

import (
	"sync"
)

const (
	DefaultSize = 5
)

type Service struct {
	size  int
	mu    sync.Mutex
	cache map[string]string
	queue Queue
}

func NewService(size int) LRUCache {
	return &Service{
		size:  size,
		cache: make(map[string]string, size),
	}
}

func (s *Service) Add(key, value string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.cache[key]
	if ok {
		return false
	}

	if s.size == 1 {
		clear(s.cache)
		s.cache[key] = value
		s.queue = Queue{}
		s.queue.add(key)
		return true
	}

	if len(s.cache) >= s.size {
		last := s.queue.removeLast()

		if last != nil {
			delete(s.cache, *last)
		}
	}

	s.queue.add(key)
	s.cache[key] = value

	return true
}

func (s *Service) Get(key string) (value string, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, ok = s.cache[key]
	if ok {
		s.queue.makeHeadIfExist(key)

		return value, true
	}

	return "", false
}

func (s *Service) Remove(key string) (ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.cache[key]; ok {
		s.queue.remove(key)
		delete(s.cache, key)
		return true
	}

	return false
}
