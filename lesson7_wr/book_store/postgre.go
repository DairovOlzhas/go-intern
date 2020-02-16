package book_store

import (
	"encoding/json"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	//"strconv"
)

type PostgreConfig struct {
	User string
	Password string
	Port string
	Host string
}

func NewPostgreBookStore(config PostgreConfig) (EndPoint, error) {
	db := pg.Connect(&pg.Options{
		Addr: config.Host + ":" + config.Port,
		User: "postgres",
		Password: config.Password,
	})
	err := createSchema(db)
	if err != nil {
		return nil, err
	}
	return &postgreStore{db: db}, nil
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Book)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

type postgreStore struct {
	db *pg.DB
}


func (ps *postgreStore) BooksCreateHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var book *Book
		err := decoder.Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Incorrect format of book"))
			return
		}
		created_book, err := ps.CreateBook(book)
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

func (ps *postgreStore) BooksListHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := ps.GetListBook()
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

func (ps *postgreStore) BookGetHandler(idParam string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars[idParam])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}
		bk, err := ps.GetBook(id)
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

func (ps *postgreStore) BookUpdateHandler(idParam string) func(http.ResponseWriter, *http.Request) {
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

		book , err = ps.UpdateBook(book, id)
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
		w.WriteHeader(201)
		w.Write(b)
	}

}

func (ps *postgreStore) BookDeleteHandler(idParam string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars[idParam])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}
		err = ps.DeleteBook(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry: " + err.Error()))
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}



func (ps *postgreStore) CreateBook(book *Book) (*Book, error) {
	return book, ps.db.Insert(&book)
}

func (ps *postgreStore) GetListBook()([]Book, error){
	var books []Book
	err := ps.db.Model(&books).Select()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (ps *postgreStore) GetBook(id int)(*Book, error){
	book := &Book{ID: id}
	err := ps.db.Select(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (ps *postgreStore) UpdateBook( book *Book, id int) (*Book, error) {
	book.ID = id
	err := ps.db.Update(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (ps *postgreStore) DeleteBook(id int) error {
	book := &Book{ID: id}
	err := ps.db.Delete(book)
	if err != nil {
		return err
	}
	return nil
}