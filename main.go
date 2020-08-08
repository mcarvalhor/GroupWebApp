// +build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, the URL path is %s!", r.URL.Path[1:])
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test 1!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/test/", handler2)
	http.HandleFunc("/", handler1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
