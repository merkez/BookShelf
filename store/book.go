package store

type Book struct {
	Name    string
	ISBN    string
	Author  string
	AddedBy string
}

func CreateBook(name, isbn, author, addedBy string) *Book {
	return &Book{
		Name:    name,
		ISBN:    isbn,
		Author:  author,
		AddedBy: addedBy,
	}
}

func (b *Book) Clone() *Book {
	return &Book{
		Name:    b.Name,
		ISBN:    b.ISBN,
		Author:  b.Author,
		AddedBy: b.AddedBy,
	}
}
