package main

import (
	"fmt"
	"golang7days/GeeWeb/Day1HttpBasic/base3/gee"
	"net/http"
)

func main() {
	engine := gee.New()
	fmt.Println("dsa")
	engine.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome")
	})
	engine.Run()

}
