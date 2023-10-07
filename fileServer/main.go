package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	route := httprouter.New()
	log.Fatal(http.ListenAndServe(":8000", route))
}
