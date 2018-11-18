package main

import (
	"fmt"
	"net/http"
	"os"
)

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	html := "Hello World"
	w.Write([]byte(html))
}

func SayHelloSasha(w http.ResponseWriter, r *http.Request) {
	html := "<b>Hello Sasha</b>"
	w.Write([]byte(html))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", SayHelloWorld)
	mux.HandleFunc("/sasha/", SayHelloSasha)

	if err := http.ListenAndServe(":555", mux); err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
}