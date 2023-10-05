package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	io.WriteString(w, currentTime.String())
}

func main() {
	http.HandleFunc("/healthcheck", HealthCheck)
	fmt.Println("Server start at port: 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
