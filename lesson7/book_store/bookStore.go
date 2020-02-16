package book_store

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type BookStore interface {
	createBook(book *Book) (*Book, error)
	listBooks() ([]*Book, error)
	getBook(id int) (*Book, error)
	updateBook(book *Book, id int) (*Book, error)
	deleteBook(id int) error
}

func CreateBookStore(path string) (*BookStoreClass, error) {
	bs := &BookStoreClass{}
	f := &os.File{}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err = os.Create(path)
		if err != nil {
			return nil, err
		}

		data, err := json.Marshal(bs)
		if err != nil {
			return nil, err
		}
		_, err = f.Write(data)
		if err != nil {
			return nil, err
		}
		err = f.Close()
		if err != nil {
			return nil, err
		}
	} else {
		f, err = os.Open(path)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		reader := bufio.NewReader(f)
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &bs)
		if err != nil {
			return nil, err
		}
	}
	return bs, nil
}

func (bs *BookStoreClass) SaveBookStore(path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	data, err := json.Marshal(bs)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (bs *BookStoreClass) createBook(book *Book) (*Book, error) {
	book.ID = bs.LastBookID
	bs.LastBookID++
	bs.Books = append(bs.Books, book)
	return book, nil
}

func (bs *BookStoreClass) listBooks() ([]*Book, error) {
	return bs.Books, nil
}

func (bs *BookStoreClass) getBook(id int) (*Book, error) {
	for _, b := range bs.Books {
		if b.ID == id {
			return b, nil
		}
	}
	return nil, errors.New("Book ID not found")

}

func (bs *BookStoreClass) updateBook(book *Book, id int) (*Book, error) {
	for i, b := range bs.Books {
		if b.ID == id {
			book.ID = id
			bs.Books[i] = book
			return bs.Books[i], nil
		}
	}
	return nil, errors.New("Book ID not found")
}

func (bs *BookStoreClass) deleteBook(id int) error {
	for i, b := range bs.Books {
		if b.ID == id {
			bs.Books = append(bs.Books[:i], bs.Books[i+1:]...)
			return nil
		}
	}
	return errors.New("Book ID not found")
}
