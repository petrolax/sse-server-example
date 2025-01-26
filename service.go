package sse

import (
	"log"
	"sync"
)

type Service struct {
	value string
	mu    sync.RWMutex
}

func NewService() *Service {
	return &Service{
		value: "default",
		mu:    sync.RWMutex{},
	}
}

func (s *Service) SetValue(newValue string) {
	log.Println("Value updated")
	s.mu.Lock()
	defer s.mu.Unlock()
	s.value = newValue
}

func (s *Service) GetValue() string {

	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.value
}
