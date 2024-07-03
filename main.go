package main

import (
	"fmt"
	"net/http"
)

func home_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "gg")
}
func contacts_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "cont")
}
func donate_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "donate")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.HandleFunc("/donate/", donate_page)

	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
