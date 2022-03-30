package controller

import (
	"encoding/json"
	"goproject/config"
	"goproject/pkg/constants"
	"goproject/pkg/service"
	"net/http"
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

	config.SetSession()
	session, _ := config.GetSession(r, constants.USER_COOKIE_NAME)

	if result != -1 {

		session.Values[constants.AUTHENTICATED_COOKIE_VALUES] = true
		session.Values[constants.USERID_COOKIE_VALUES] = result

		session.Save(r, w)

		resp["redirect"] = string(constants.DASHBOARD_ROUTE)
		w.WriteHeader(http.StatusAccepted)

	} else {

		session.Values["authenticated"] = false
		session.Values["userid"] = -1.

		resp["error"] = "Username or password is incorrect"
		w.WriteHeader(http.StatusUnauthorized)
	}

	session.Save(r, w)

	json.NewEncoder(w).Encode(resp)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	username := r.Form["regi_username"][0]
	name := r.Form["regi_name"][0]
	password := r.Form["regi_password"][0]

	var result bool

	// fmt.Println("\n\nform un: " + username[0])
	// fmt.Println("\n\nform name: " + name[0])
	// fmt.Println("\n\nform " + password[0])

	result = service.CheckIsNewUserName(username)

	// fmt.Println("\n\nNew user: ", result, "\n\n")

	resp := make(map[string]string)

	w.Header().Set("Content-Type", "application/json")

	if result {
		result = service.CreateUser(username, password, name)

		if result {
			resp["message"] = "Registration successful"
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			resp["error"] = "Something went wrong! unable to register"
		}

	}

	if result {

		w.WriteHeader(http.StatusAccepted)

	} else {

		resp["error"] = "Username is already in use Please choose again"
		w.WriteHeader(http.StatusUnauthorized)
	}

	json.NewEncoder(w).Encode(resp)

}
