package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)


type Book struct {
        ID int
        ISBN string
        Author string
        PublishedYear string
}

func LogHandle(w http.ResponseWriter, r *http.Request)  {
    log.Printf("%q", r.UserAgent())

    book := Book{
        ID: 12,
        ISBN: "hIbUI-92JL-k13l",
        Author: "amsem",
        PublishedYear: "1999",
    }
    jsonData, _ := json.Marshal(book)
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
}

func main()  {
    f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("error whern opening file %v", err)
    }
    defer f.Close()
    log.SetOutput(f)
    http.HandleFunc("/", LogHandle)
    s := &http.Server{
        Addr: ":9090",
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    log.Fatal(s.ListenAndServe())
}
