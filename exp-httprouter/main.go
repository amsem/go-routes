package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func getCommandOutput(command string, argumets ...string) string {
	out, _ := exec.Command(command, argumets...).Output()
	return string(out)
}

func goVersion(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response := getCommandOutput("/usr/local/go/bin/go", "version")
	io.WriteString(w, response)
	return
}

func getFileContent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, getCommandOutput("/bin/cat", p.ByName("name")))
}

func main() {
	router := httprouter.New()

	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)

	log.Fatal(http.ListenAndServe(":8000", router))
}
