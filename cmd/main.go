package main

import (
	"fmt"
	"goproject/config"
	"goproject/pkg/db"
	"goproject/pkg/routes"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	r := routes.InitializeHandlers()

	openbrowser()

	config.UnsetSession()

	http.ListenAndServe(":8080", r)

	db.CloseConnection()

	println("Connection closed")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func openbrowser() {

	url := "http://localhost:8080/login"

	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
