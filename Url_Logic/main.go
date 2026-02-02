package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	})

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	fmt.Println("Listening on port 3000")
	server.ListenAndServe()

}
