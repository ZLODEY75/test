package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                 string
	Age                  uint16
	Money                int16
	Avg_grads, Happiness float64
	Hoobbies             []string
}

func (u *User) getAllinfo() string {
	return fmt.Sprintf("User name is: %s He is %d and he has meney: %d", u.Name, u.Age, u.Money)
}
func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"bob", 25, -50, 4.2, 0.8, []string{"Footbboll", "Skate", "Danse"}}
	//bob.setNewName("Alex")
	//fmt.Fprintf(page, "<b>Main text<b>")
	tmpl, _ := template.ParseFiles("templetes/home_page.html")
	tmpl.Execute(w, bob)

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
