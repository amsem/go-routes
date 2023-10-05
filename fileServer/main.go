package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	route := httprouter.New()
	route.ServeFiles("/static/*filepath", http.Dir("/home/amsem/workspace/src/github.com/amsem/go-routes/fileServer/static"))
	log.Fatal(http.ListenAndServe(":8000", route))
}
