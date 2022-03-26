package controller

import (
	"fmt"
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

func renderTemplateWithData(w http.ResponseWriter, page string, data []byte) {
	// func renderTemplate(params ...interface{}) {

	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Println(err)
		return
	}
}
