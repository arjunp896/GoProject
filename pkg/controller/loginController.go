package controller

import (
	"encoding/json"
	"goproject/pkg/service"
	"net/http"

	"github.com/gorilla/mux"
)

func ValidateUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	username := r.Form["username"]
	password := r.Form["password"]

	// fmt.Println(username)
	// fmt.Println(password)

	result := service.ValidateUser(username[0], password[0])

	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)

	if result {

		resp["redirect"] = "/car/cars/"
		w.WriteHeader(http.StatusAccepted)

	} else {

		resp["error"] = "Username or password is incorrect"
		w.WriteHeader(http.StatusUnauthorized)
	}

	json.NewEncoder(w).Encode(resp)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	username := params["username"]
	name := params["name"]
	password := params["password"]

	var result bool

	result = service.CheckIsNewUserName(username)

	resp := make(map[string]string)

	w.Header().Set("Content-Type", "application/json")

	if result {
		result = service.CreateUser(username, password, name)

		if result {
			w.WriteHeader(http.StatusCreated)
			resp["message"] = "Registration successful"
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			resp["error"] = "Something went wrong! unable to register"
		}

	}

	if result {

		w.WriteHeader(http.StatusAccepted)

	} else {

		resp["error"] = "Username or password is incorrect"
		w.WriteHeader(http.StatusUnauthorized)
	}

	json.NewEncoder(w).Encode(resp)

}
