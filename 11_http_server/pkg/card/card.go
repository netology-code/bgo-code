package card

import (
	"context"
	"sync"
)

type Card struct {
	Id int64
	Number string
}

type Service struct {
	mu sync.RWMutex
	cards []*Card
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) All(context.Context) []*Card {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cards
}


