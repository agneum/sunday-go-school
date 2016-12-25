package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error. Bad request\n")
		// http.Error(w, "Error. Bad request", 400)
		return
	}

	fmt.Fprintf(w, "Hello %s\n", name)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
