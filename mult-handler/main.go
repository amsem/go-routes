package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	newMux := http.NewServeMux()
	newMux.HandleFunc("/randFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})
	newMux.HandleFunc("/randInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Int())
	})
	fmt.Println("SERVER PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", newMux))
}
