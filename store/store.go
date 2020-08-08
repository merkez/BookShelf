package store

import (
	"errors"
	"sync"
)

var (
	ErrBookAlreadyExist = errors.New("Book is already exists on BookShelf")
	ErrBookNotFound     = errors.New("Book is NOT found on BookShelf")
)

type BookStore interface {
	Add(b *Book) error
	List() []*Book
	Del(isbn string) (string, error)
	Find(isbn string) (*Book, error)
}

type InMemoryStore struct {
	m     sync.Mutex
	books map[string]*Book
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{books: make(map[string]*Book)}
}

func (store *InMemoryStore) Del(isbn string) (string, error) {
	store.m.Lock()
	defer store.m.Unlock()
	book := store.books[isbn]
	if book == nil {
		return "Error: ", ErrBookNotFound
	}
	delete(store.books, isbn)
	return "Book with ISBN " + isbn + " deleted successfully", nil
}

func (store *InMemoryStore) Find(isbn string) (*Book, error) {
	store.m.Lock()
	defer store.m.Unlock()
	book := store.books[isbn]
	if book == nil {
		return nil, ErrBookNotFound
	}
	return book, nil
}

func (store *InMemoryStore) Add(book *Book) error {
	store.m.Lock()
	defer store.m.Unlock()
	if store.books[book.ISBN] != nil {
		return ErrBookAlreadyExist
	}
	store.books[book.ISBN] = book.Clone()
	return nil
}

func (store *InMemoryStore) List() []*Book {
	store.m.Lock()
	defer store.m.Unlock()
	var books []*Book
	for _, v := range store.books {
		books = append(books, v)
	}
	return books
}
