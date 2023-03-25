package main

import (
	"fmt"
	"log"
	"net/http"
)

type myHandler struct {
}

func (handler *myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "url.path = %q\n", req.URL.Path)
	}
}

func main() {
	myHandler := new(myHandler)
	log.Fatal(http.ListenAndServe(":9999", myHandler))
}
