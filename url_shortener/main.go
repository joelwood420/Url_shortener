package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type ShorterRequest struct {
	Url string `json:"Url"`
	Tag string `json:"Tag"`
}

type ShorterResponse struct {
	ShortURL string `json:"short_url"`
}

var urlMap = make(map[string]string)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("URL Shortener API"))
	})

	router.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {
		var requestData ShorterRequest
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		shortCode := generateShortURL(requestData.Url)
		urlMap[shortCode] = requestData.Url

		response := ShorterResponse{ShortURL: fmt.Sprintf("https://url-shortener-damp-pond-9136.fly.dev/%s", shortCode)}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	router.HandleFunc("GET /{shortCode}", func(w http.ResponseWriter, r *http.Request) {
		shortCode := r.PathValue("shortCode")
		longURL, exists := urlMap[shortCode]
		if !exists {
			http.Error(w, "Short URL not found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, longURL, http.StatusFound)
	})

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	fmt.Println("URL Shortener listening on port 3000")
	server.ListenAndServe()
}

func generateShortURL(longURL string) string {
	id := uuid.New()
	shortCode := strings.ReplaceAll(id.String()[:8], "-", "")
	return shortCode
}
