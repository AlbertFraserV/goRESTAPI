package main

import "fmt"
import "net/http"
import "log"

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }


    fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request successful")
    name := r.FormValue("name")
    address := r.FormValue("address")

    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
    fileServer := http.FileServer(http.Dir("./static")) // New code
    http.Handle("/", fileServer)

    fmt.Printf("Starting server. Port 8080.\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

