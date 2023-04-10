package main

import (
	"bookstore/book"
	"bookstore/database"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

var books []book.Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	books = database.GetBooks()
	w.Header().Set("Content-Type", "application/json")
	base := r.URL.Query().Get("filter")
	//fmt.Println(base)
	if base == "asc" {
		asc := database.FilterByPriceAsc()
		json.NewEncoder(w).Encode(asc)
	} else if base == "desc" {
		desc := database.FilterByPriceDesc()
		json.NewEncoder(w).Encode(desc)
	} else {
		json.NewEncoder(w).Encode(books)
	}

	fmt.Println("GetBooks Endpoint")

}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	value := r.URL.Query().Get("id")
	for _, item := range books {
		if item.ID == value {
			json.NewEncoder(w).Encode(item)
			fmt.Println(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&book.Book{})
	fmt.Println("GetBooksByID Endpoint")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book book.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	name := r.URL.Query().Get("title")
	descriprion := r.URL.Query().Get("description")
	price := r.URL.Query().Get("price")
	book.ID = strconv.Itoa(rand.Intn(10000000))
	book.Title = name
	book.Description = descriprion
	marks, err := strconv.Atoi(price)
	if err != nil {
		fmt.Println("Error with price")
		return
	}
	book.Price = marks
	books = append(books, book)
	database.Insert(book)
	json.NewEncoder(w).Encode(book)
	fmt.Println("POST Endpoint")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("title")
	descriprion := r.URL.Query().Get("description")
	price := r.URL.Query().Get("price")
	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			var book book.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = id
			book.Title = name
			book.Description = descriprion
			marks, err := strconv.Atoi(price)
			if err != nil {
				fmt.Println("Error with price")
				return
			}
			book.Price = marks
			database.UpdateById(book)
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
	fmt.Println("PUT Endpoint")
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	value := r.URL.Query().Get("id")
	for index, item := range books {
		if item.ID == value {
			books = append(books[:index], books[index+1:]...)
			database.DeleteById(item.ID)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
	fmt.Println("DELETE Endpoint")
}

func main_page(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	for k, v := range values {
		fmt.Println(k, " => ", v)
	}
}

func searchBook(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query().Get("title")
	values = "%" + values + "%"
	book := database.SearchByName(values)
	json.NewEncoder(w).Encode(book)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", main_page)
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book", getBook).Methods("GET")
	r.HandleFunc("/book", createBook).Methods("POST")
	r.HandleFunc("/book", updateBook).Methods("PUT")
	r.HandleFunc("/book", deleteBook).Methods("DELETE")
	r.HandleFunc("/search", searchBook).Methods("GET")
    fmt.Println("server is started")
	http.ListenAndServe(":3000", r)
}
