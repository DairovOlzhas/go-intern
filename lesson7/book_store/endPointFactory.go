package book_store

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EndPoint interface {
	BookGetHandler(idParam string) func(http.ResponseWriter, *http.Request)
	BookUpdateHandler(idParam string) func(http.ResponseWriter, *http.Request)
	BooksListHandler() func(http.ResponseWriter, *http.Request)
	BooksCreateHandler() func(http.ResponseWriter, *http.Request)
	BookDeleteHandler(idParam string) func(http.ResponseWriter, *http.Request)
}

type endPointFactory struct {
	bks *BookStoreClass
}

func CreateEndPointFactory(bookStore *BookStoreClass) (EndPoint, error) {
	return &endPointFactory{bks: bookStore}, nil
}

func SaveBookStoreHandler(bookStore *BookStoreClass, pathToBooksStore string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := bookStore.SaveBookStore(pathToBooksStore)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Book Store saved successfully!"))
	}
}

func (ef *endPointFactory) BooksCreateHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var book *Book
		err := decoder.Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Incorrect format of book"))
			return
		}
		created_book, err := ef.bks.createBook(book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}
		data, err := json.Marshal(created_book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	}
}

func (ef *endPointFactory) BooksListHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := ef.bks.listBooks()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}
		n, err := json.Marshal(books)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(""))
			return
		}
		w.WriteHeader(200)
		w.Write(n)
	}

}

func (ef *endPointFactory) BookGetHandler(idParam string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars[idParam])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}
		bk, err := ef.bks.getBook(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}
		book, err := json.Marshal(bk)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(book)
	}
}

func (ef *endPointFactory) BookUpdateHandler(idParam string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars[idParam])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}

		decoder := json.NewDecoder(r.Body)

		var book *Book
		err = decoder.Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Incorrect format of book"))
			return
		}

		book, err = ef.bks.updateBook(book, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}
		b, err := json.Marshal(book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}
		w.WriteHeader(200)
		w.Write(b)
	}
}

func (ef *endPointFactory) BookDeleteHandler(idParam string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars[idParam])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}
		err = ef.bks.deleteBook(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
