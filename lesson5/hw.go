package main

import (
	"bufio"
	"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	s "strings"
)

var (
	host = "8080"
	id = 0
)




type Book struct {
	//id int
	Name string `json:"name"`
	Description string `json:"description"`
	Author string `json:"author"`
}

func createBook(book Book){
	n, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}

	_, err = os.Open("books/"+strconv.Itoa(id)+".json")
	for err == nil {
		id++
		_, err = os.Open("books/"+strconv.Itoa(id)+".json")
	}

	fo, err := os.Create("books/"+strconv.Itoa(id)+".json")

	if err != nil {
		panic(err)
	}

	_, err = fo.Write(n)
	if err != nil {
		panic(err)
	}

	err = fo.Close()
	if err != nil {
		panic(err)
	}
}

func getBook(id int) Book {
	a := Book{}
	f, err := os.Open("books/"+strconv.Itoa(id)+".json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	file, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &a)
	if err != nil {
		panic(err)
	}

	return a
}


func updateBook(book Book, lid int) []byte {
	deleteBook(lid)

	n, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}


	fo, err := os.Open("books/"+strconv.Itoa(lid)+".json")
	if err != nil {
		panic(err)
	}
	fo.Truncate(0)
	_, err = fo.Write(n)
	if err != nil {
		panic(err)
	}

	err = fo.Close()
	if err != nil {
		panic(err)
	}
	return n
}

func deleteBook(id int){
	err := os.Remove("books/"+strconv.Itoa(id)+".json")
	if err != nil {
		panic(err)
	}
}

func booksCreateHandler(w http.ResponseWriter, r *http.Request){

	decoder := json.NewDecoder(r.Body)
	var book Book
	err := decoder.Decode(&book)
	if err != nil {
		panic(err)
	}
	createBook(book)
	w.WriteHeader(201)
}

func booksListHandler(w http.ResponseWriter, r *http.Request){
	books, err := ioutil.ReadDir("books")
	if err != nil {
		log.Fatal(err)
	}
	bks := make([]Book, 0)
	for i:=0; i < len(books); i++ {
		lid, _ := strconv.Atoi(s.Split(books[i].Name(),".")[0])
		bks = append(bks, getBook(lid))
		//fmt.Println(s.Split(books[i].Name(),".")[0])
	}
	n, err := json.Marshal(bks)

	if err != nil {
		panic(err)
	}

	w.Write(n)

}

func bookGetHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	lid, _ := strconv.Atoi(vars["id"])

	n, err := json.Marshal(getBook(lid))
	if err != nil {
		panic(err)
	}

	w.Write(n)
}

func bookUpdateHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	lid, _ := strconv.Atoi(vars["id"])
	decoder := json.NewDecoder(r.Body)
	var book Book
	err := decoder.Decode(&book)
	if err != nil {
		panic(err)
	}
	w.Write(updateBook(book, lid))
	w.WriteHeader(202)

}

func bookDeleteHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	deleteBook(id)
	w.WriteHeader(204)

}

func main(){
	//book :=  Book{
	//	Name:        "coladd",
	//	Description: "a asdf of pades",
	//	Author:      "loram",
	//}
	//
	//createBook(book)

	r := mux.NewRouter()
	r.HandleFunc("/books", booksListHandler).Methods("GET")
	r.HandleFunc("/books", booksCreateHandler).Methods("POST")
	r.HandleFunc("/book/{id:[0-9]+}", bookGetHandler).Methods("GET")
	r.HandleFunc("/book/{id:[0-9]+}", bookUpdateHandler).Methods("PUT")
	r.HandleFunc("/book/{id:[0-9]+}", bookDeleteHandler).Methods("DELETE")

	http.ListenAndServe(":"+host, r)

}
