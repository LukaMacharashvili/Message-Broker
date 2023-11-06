package main

// TODO: Dockerize
import (
	"net/http"

	"github.com/LukaMacharashvili/Message-Broker/internal/handlers"
)

func main() {
	handlers := handlers.Handlers{
		ConsumersMap: make(map[string][]string),
	}

	http.HandleFunc("/register", handlers.Register)

	http.HandleFunc("/publish", handlers.Publish)

	http.ListenAndServe(":3003", nil)
}
