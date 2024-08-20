package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sesi-lima/helpers"
)

var PORT = ":3000"

func main() {
	http.HandleFunc("/", loginPage)
	http.HandleFunc("/login", loginProcess)

	fmt.Println("Server Running on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tpl.Execute(w, helpers.People)
		return
	}
}
func loginProcess(w http.ResponseWriter, r *http.Request) {
	condition := false
	r.ParseForm()
	var peoples helpers.Peoples
	for i, people := range helpers.People {
		if people.Email == r.FormValue("email") && people.Password == r.FormValue("password") {
			condition = true
			peoples = helpers.People[i]
			break
		}
	}
	if !condition {
		tpl, err := template.ParseFiles("not-registered.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tpl.Execute(w, nil)
		return
	}
	tpl, err := template.ParseFiles("detail.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, peoples)
}
