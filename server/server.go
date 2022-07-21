package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var endPoints = [...]string{"book"}

func checkEndPoints(endPointRequest string) bool {
	for i := range endPoints {
		if endPoints[i] == endPointRequest {
			return true
		}
	}
	return false
}

type book struct {
	ID     string `json:"ID"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
}

type allBooks []book

var events = allBooks{
	{
		ID:     "1",
		Title:  "Slaughterhouse-Five",
		Author: "Vonnegutt",
	},
	{
		ID:     "2",
		Title:  "The Lightning Thief",
		Author: "Riordan",
	},
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !checkEndPoints(strings.Trim(r.URL.Path, "/")) {
		http.Error(w, r.URL.Path, http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(events)
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
	for i := range endPoints {
		http.HandleFunc("/"+endPoints[i], helloHandler)
	}
	http.HandleFunc("/form", formHandler)
	fileServer := http.FileServer(http.Dir("./static")) // New code
	http.Handle("/", fileServer)

	fmt.Printf("Starting server. Port .\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
