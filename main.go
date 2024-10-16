package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type BookType struct { 
    Title string `json:"title"`
    Author string `json:"author"`
    Pages int `json:"pages"`
}

var Books = map[string]BookType {
	"War and Peace": {
		Title: "War and Peace",
        Author: "Leo Tolstoy",
        Pages: 1225,
	},
	"The Great Gatsby": {
		Title: "The Great Gatsby",
        Author: "F. Scott Fitzgerald",
        Pages: 180,
	},
	"The Catcher in the Rye": {
		Title: "The Catcher in the Rye",
		Author: "J.D. Salinger",
		Pages: 234,
	},
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    title := vars["title"]
    json.NewEncoder(w).Encode(Books[title])
}

func GetAllBooksByAuthor(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    author := vars["author"]
    
    var authorBooks []BookType
    for _, book := range Books {
        if book.Author == author {
            authorBooks = append(authorBooks, book)
        }
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(authorBooks)
}
    
func main() {
    r := mux.NewRouter()

    r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })

    // ** Features of the gorilla/mux Router**
    
    // Path Prefixes & Subrouters
    // Restrict the request handler to specific path prefixes.
	bookrouter := r.PathPrefix("/books").Subrouter()

    // Hostnames & Subdomains
    // Restrict the request handler to specific hostnames or subdomains.
	bookrouter.HandleFunc("/", AllBooks).Host("localhost:8000")

    // Schemes
    // Restrict the request handler to http/https.
	bookrouter.HandleFunc("/{title}", GetBook).Schemes("http")

    // Methods
    // Restrict the request handler to specific HTTP methods.
	bookrouter.HandleFunc("/author/{author}", GetAllBooksByAuthor).Methods("GET")

    http.ListenAndServe(":8000", r)
}