package handlers

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
)

func SayHello(w http.ResponseWriter, name string) {

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, http.StatusText(http.StatusBadRequest))
		logrus.Error("The 'name' parameter is missed")
		return
	}

	fmt.Fprintf(w, "Hello %s\n", name)
	logrus.Infof("Handling a request with the name: %s", name)
}
