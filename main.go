package main

import (
	"net/http"
	"fmt"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sorry %s is the wrong answer", r.URL.Path)
	})
	http.HandleFunc("/question", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "What is 2 plus 2", r.URL.Path)
	})
	http.HandleFunc("/4", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "4 is correct", r.URL.Path)
	})
}
