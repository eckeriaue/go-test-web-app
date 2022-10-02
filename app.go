package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "second test")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contact page")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactPage)
	http.ListenAndServe(":1111", nil)
}

func main() {
	handleRequest()
}
