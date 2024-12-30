package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/kyleborot/golearn/url-shortener/models"
	"google.golang.org/api/option"
)

var Client *firestore.Client

func InitializeFirestore() (*firestore.Client, error) {
	ctx := context.Background()
	sakPath := "C:/Users/kborot/Desktop/golearn/url-shortener/serviceAccountKey.json"
	client, err := firestore.NewClient(ctx, "url-shortener-446004", option.WithCredentialsFile(sakPath))
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	Client = client
	return client, nil
}
func StoreShortenedURL(ctx context.Context, urlData models.UrlStruct) error {
	collectionRef := Client.Collection("urls")
	_, err := collectionRef.Doc(urlData.ID).Set(ctx, map[string]interface{}{
		"long_url":  urlData.LongURL,
		"timestamp": urlData.Timestamp,
	})
	if err != nil {
		log.Printf("Failed to store URL: %v", err)
		return err
	}
	return nil
}
