package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ShorterRequest struct {
	Url string
	Tag string
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	})

	router.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {

		requestData := ShorterRequest{}
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			fmt.Println("JSON decode error:", err) //debug
			panic(err)
		}

		w.Write([]byte(requestData.Url))
	})

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	fmt.Println("Listening on port 3000")
	server.ListenAndServe()

}
