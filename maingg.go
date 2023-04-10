package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bookstore/book"
	"github.com/gorilla/mux"
	"strconv"
	"math/rand"
	"bookstore/database"
)


var Books []book.Book

func main_page(w http.ResponseWriter, r *http.Request){
	fmt.Println("Hello world")
}

func getAllBooks(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(Books)
	fmt.Println("GetBooks Endpoint")
}

func getBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    key, _ := strconv.ParseInt(vars["id"], 10, 64)
    for _, book := range Books {
        if book.Id == key {
            json.NewEncoder(w).Encode(book)
        }
    }
	fmt.Println("GetBookByID Endpoint")
}

func UpdateBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    key, _ := strconv.ParseInt(vars["id"], 10, 64)
    for _, book := range Books {
        if book.Id == key {
            json.NewEncoder(w).Encode(book)
        }
    }
	fmt.Println("PUT Endpoint")
}

func InsertBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    var book book.Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.Id = int64(rand.Intn(1000000))
	fmt.Println(book)
    json.NewEncoder(w).Encode(book)
	fmt.Println("POST Endpoint")
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    key, _ := strconv.ParseInt(vars["id"], 10, 64)
    for _, book := range Books {
        if book.Id == key {
            database.DeleteById(key)
			Books = database.GetBooks()
        }
    }
	fmt.Println("DELETE Endpoint")
}

func handleFunctions(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", main_page)
	myRouter.HandleFunc("/books", getAllBooks).Methods("GET")
	myRouter.HandleFunc("/books/{id}", getBookById).Methods("GET")
	myRouter.HandleFunc("/books/{id}", UpdateBookById).Methods("PUT")
	myRouter.HandleFunc("/books", InsertBook).Methods("POST")
	myRouter.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
	http.ListenAndServe(":2004", myRouter)
	fmt.Println("Server is running")
	
}

func maingg(){
	Books = database.GetBooks()
	//fmt.Print(Books)
	handleFunctions()
	//database.ConnectDB()
	//fmt.Println(database.GetBooks())
	//database.UpdateById(3, "TestName", "TestDescription")
	//
	//fmt.Println(database.GetById(5))
	//database.Insert(Books[0])
	//fmt.Println(database.SearchByName(Books[0].Name))
	//fmt.Println(database.FilterByPriceDesc())
}
