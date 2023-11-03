package main

// TODO: Concurrency when modifying a map
// TODO: separate functionality into packages
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getRequestBody(r *http.Request) []byte {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	return body
}

var consumersMap map[string][]string

func main() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
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

		if consumersMap == nil {
			consumersMap = make(map[string][]string)
		}

		consumersMap[topic] = append(consumersMap[topic], consumer+handlerPath)
	})

	http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		topic := r.URL.Query().Get("topic")

		if topic == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("topic is required"))
			return
		}

		fmt.Printf("Topic: %s\n", topic)
		consumers := consumersMap[topic]
		fmt.Printf("Consumers: %v\n", consumers)
		for _, consumer := range consumers {
			go func(consumer string, body []byte) {
				_, err := http.Post(consumer, "application/json", bytes.NewBuffer(body))
				if err != nil {
					// TODO: Send to DLQ
					return
				}
			}(consumer, getRequestBody(r))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Message processing started"))
	})

	http.ListenAndServe(":3003", nil)
}
