package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		log.Println("404 Not Found at", r.RequestURI)
		return
	}

	w.Write([]byte("Hello From SNIPPETU"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allowed", http.MethodPost)
		http.Error(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create new snippet"))
}

func googleIt(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://google.com", http.StatusFound)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/google/", googleIt)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Server starting at http://localhost:4000/")
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}
