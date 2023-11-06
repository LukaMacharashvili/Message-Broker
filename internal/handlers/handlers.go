package handlers

import "sync"

type Handlers struct {
	ConsumersMap      map[string][]string
	ConsumersMapMutex sync.RWMutex
}
