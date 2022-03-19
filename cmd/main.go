package main

import (
	"fmt"
	"goproject/pkg/routes"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {

	r := routes.InitializeHandlers()

	openbrowser()

	http.ListenAndServe(":8080", r)
}

func openbrowser() {

	url := "http://localhost:8080"

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
