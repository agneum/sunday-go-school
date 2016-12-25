package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error. Bad request\n")
		return
	}

	fmt.Fprintf(w, "Hello %s\n", name)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error. Bad request\n")
	}

	fmt.Fprintf(w, "Hello %s!\n", body)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/hello", Hello)

	log.Fatal(http.ListenAndServe(":8000", router))
}
