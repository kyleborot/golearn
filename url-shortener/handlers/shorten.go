package handlers

import (
	"errors"
	"net/http"
	"net/url"
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

func postRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
