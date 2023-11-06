package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	h.ConsumersMapMutex.Lock()
	defer h.ConsumersMapMutex.Unlock()

	topic := r.URL.Query().Get("topic")

	if topic == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("topic is required"))
		return
	}
	fmt.Printf("Topic: %s\n", topic)

	consumer := r.Header.Get("X-Consumer")

	if consumer == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Uknown consumer"))
		return
	}
	fmt.Printf("Consumer: %s\n", consumer)

	handlerPath := r.Header.Get("X-Handler-Path")

	if handlerPath == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Handler path is required"))
		return
	}
	fmt.Printf("Handler path: %s\n", handlerPath)

	if h.ConsumersMap == nil {
		h.ConsumersMap = make(map[string][]string)
	}

	h.ConsumersMap[topic] = append(h.ConsumersMap[topic], consumer+handlerPath)
}
