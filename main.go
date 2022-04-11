package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

var restaurants = []string{"Cleo's", "Dragos", "Magazine Pizza"}

func main() {
	mux := chi.NewRouter()
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome To TableFinder!"))
	})

	go http.ListenAndServe(":3000", mux)

	fmt.Println("Listening on localhost:3000...")
	for {

	}
}
