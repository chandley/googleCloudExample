package main

import (
	"net/http"
	"fmt"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You requested %s/n", r.URL.Path)
	})
}
