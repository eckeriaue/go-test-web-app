package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello wrld")

}

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":1111", nil)
}
