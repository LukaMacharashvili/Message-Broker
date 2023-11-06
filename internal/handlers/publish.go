package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/LukaMacharashvili/Message-Broker/internal/utils/http_utils"
)

func (h *Handlers) Publish(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.Query().Get("topic")

	if topic == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("topic is required"))
		return
	}

	fmt.Printf("Topic: %s\n", topic)
	consumers := h.ConsumersMap[topic]
	fmt.Printf("Consumers: %v\n", consumers)
	for _, consumer := range consumers {
		go func(consumer string, body []byte) {
			_, err := http.Post(consumer, "application/json", bytes.NewBuffer(body))
			if err != nil {
				// TODO: Send to DLQ
				return
			}
		}(consumer, http_utils.GetRequestBody(r))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message processing started"))
}
