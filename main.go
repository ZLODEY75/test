package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                 string
	age                  uint16
	money                int16
	avg_grads, happiness float64
}

func (u *User) getAllinfo() string {
	return fmt.Sprintf("User name is: %s He is %d and he has meney: %d", u.name, u.age, u.money)
}
func (u *User) setNewName(newName string) {
	u.name = newName
}

func home_page(page http.ResponseWriter, r *http.Request) {
	bob := User{"bob", 25, -50, 4.2, 0.8}
	bob.setNewName("Alex")
	fmt.Fprintf(page, bob.getAllinfo())

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
	//bob := User{name: "bob", age: 25, money: -50, avg_grads: 4.2, happiness: 0.8}

	handleRequest()
}
