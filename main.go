package main

import (
    "io"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

func main() {

    r := mux.NewRouter()

	r.HandleFunc("/", ExampleHandler)

	fmt.Println("Server listening!")
	http.ListenAndServe(":8080", r)
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    io.WriteString(w, `{"status":"ok"}`)
}