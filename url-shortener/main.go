package main

import (
	"log"
	"net/http"

	"github.com/kyleborot/golearn/url-shortener/db"
	"github.com/kyleborot/golearn/url-shortener/handlers"
)

func main() {
	_, err := db.InitializeFirestore()
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}

	http.HandleFunc("/shorten-url", handlers.PostRequest)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
