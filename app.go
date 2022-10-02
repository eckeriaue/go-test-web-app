package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                 string
	age                  uint16
	money                int16
	avgGrades, happiness float64
}

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
	// var bob User = User{}
	//	bob := User{
	//	name:      "bob",
	//	age:       25,
	//		money:     -50,
	//		avgGrades: 4.2,
	//		happiness: 0.8,
	//	}
	bob := User{
		"bob",
		25,
		-50,
		4.2,
		0.8,
	}
	fmt.Println(bob)
	handleRequest()
}
