package handlers

import (
	"context"
	"crypto/sha256"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/kyleborot/golearn/url-shortener/db"
	"github.com/kyleborot/golearn/url-shortener/models"
	"github.com/kyleborot/golearn/url-shortener/utils"
)

func validateUrl(longURL string) (bool, error) {
	_, err := url.ParseRequestURI(longURL)
	if err != nil {
		return false, errors.New("invalid URL format")
	}

	resp, err := http.Head(longURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false, errors.New("URL Inaccessible")
	}

	return true, nil
}

func createJSONPackage(longURL string) (models.UrlStruct, error) {
	isValid, err := validateUrl(longURL)
	if !isValid || err != nil {
		return models.UrlStruct{}, errors.New("invalid or inaccessible URL" + err.Error())
	}
	hash := sha256.New()
	hash.Write([]byte(longURL))
	hashBytes := hash.Sum(nil)

	shortCode := utils.ToBase62(hashBytes[:8])

	urlEntry := models.UrlStruct{
		ID:        shortCode,
		LongURL:   longURL,
		Timestamp: time.Now(),
	}
	return urlEntry, nil
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	longURL := r.FormValue("longURL")
	if longURL == "" {
		http.Error(w, "Missing longURL parameter", http.StatusBadRequest)
		return
	}

	jsonResponse, err := createJSONPackage(longURL)
	if err != nil {
		http.Error(w, "Internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	err = db.StoreShortenedURL(ctx, jsonResponse)
	if err != nil {
		http.Error(w, "Failed to store URL", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "URL successfully shortened!"}`))
}
