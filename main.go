package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Home page of the test application")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "The about page of the test application")
	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
}
