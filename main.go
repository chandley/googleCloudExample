// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Sample helloworld is a basic App Engine flexible app.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/answer", answerHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Printf("%v",r.Form)
	if r.Form.Get("answer") != "4" {
		fmt.Fprint(w, "Wrong")
		return
	}
	fmt.Fprint(w, "Right")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
