package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
)

type UUID struct {
}

func (p *UUID) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		giveranduuid(w, req)
		return
	}
	http.NotFound(w, req)
	return
}

func giveranduuid(w http.ResponseWriter, r *http.Request) {
	c := 900
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, fmt.Sprintf("%x", b))
}

func main() {
	mux := &UUID{}
	fmt.Println("Server Running on PORT : 8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
