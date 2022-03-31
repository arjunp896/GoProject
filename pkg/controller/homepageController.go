package controller

import (
	"fmt"
	"goproject/config"
	"goproject/pkg/constants"
	"html/template"
	"log"
	"net/http"
)

func Index_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Jai Swaminarayan from goproject")
}

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {

	// fmt.Print("\nload login page")

	renderTemplate(w, "./web/login.html")

}

func LoadDashboardPage(w http.ResponseWriter, r *http.Request) {

	session, _ := config.GetSession(r, constants.USER_COOKIE_NAME)

	if _, ok := session.Values[constants.AUTHENTICATED_COOKIE_VALUES].(bool); ok {
		renderTemplate(w, "./web/dashboard.html")
	} else {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
	}

}

func LoadMyCars(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "./web/vehicles.html")

}

func renderTemplate(w http.ResponseWriter, page string) {
	// func renderTemplate(params ...interface{}) {

	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func renderTemplateWithData(w http.ResponseWriter, page string, data interface{}) {
	// func renderTemplate(params ...interface{}) {

	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println("render with data parse in template", err)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Println(err)
		return
	}
}
