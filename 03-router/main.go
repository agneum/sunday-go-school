package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/agneum/sunday-go-school/03-router/handlers"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := r.URL.Query().Get("name")
	handlers.SayHello(w, name)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusText(http.StatusInternalServerError))
	}

	handlers.SayHello(w, string(body))
}

func main() {
	logrus.Infoln("Server is running..")

	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/hello", Hello)

	logrus.Fatal(http.ListenAndServe(":8000", router))
}
