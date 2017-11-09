package main

import (
	"net/http"
	"fmt"
)

func Main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You requested %s/n", r.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}
