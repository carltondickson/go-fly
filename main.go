package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	var err error
	var files *template.Template
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	files, err = template.ParseFiles("./templates/index.html")
	if err != nil {
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}
		err = files.Execute(w, data)
		if err != nil {
			_, _ = fmt.Fprint(w, err)
			return
		}
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "The about page of the test application")
	})

	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		return
	}
}
