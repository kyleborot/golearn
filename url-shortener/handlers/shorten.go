package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

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

func createJSONPackage(longURL string) ([]byte, error) {
	isValid, err := validateUrl(longURL)
	if !isValid || err != nil {
		return nil, errors.New("invalid or inaccessible URL")
	}
	shortCode, err := utils.FromBase62(longURL)
	if err != nil {
		return nil, errors.New("could not encode URL")
	}
	urlEntry := models.UrlStruct{
		ID:        string(rune(shortCode)),
		LongURL:   longURL,
		Timestamp: time.Now(),
	}
	jsonResponse, err := json.Marshal(urlEntry)
	if err != nil {
		return nil, errors.New("failed to encode JSON")
	}
	return jsonResponse, nil
}

func postRequest(w http.ResponseWriter, r *http.Request, longURL string) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	jsonResponse, err := createJSONPackage(longURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error Occurred (failed JSONResponse)"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
