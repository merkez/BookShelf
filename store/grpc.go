package store

import (
	"context"

	pb "github.com/mrturkmencom/bookshelf/proto"
)

type BookShelfServer struct {
	B BookStore
}

func (b *BookShelfServer) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	if req.Book.Isbn == "" {
		return &pb.AddBookResponse{Message: "without ISBN number, book cannot be added"}, nil
	}
	book := CreateBook(req.Book.Name, req.Book.Isbn, req.Book.Author, req.Book.AddedBy)
	if err := b.B.Add(book); err != nil {
		return &pb.AddBookResponse{Message: "Error " + err.Error()}, err
	}
	return &pb.AddBookResponse{Message: "Book " + book.Name + " added successfully"}, nil
}

func (b *BookShelfServer) ListBook(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	var bookInfoResp []*pb.ListBooksResponse_BookInfo
	books := b.B.List()
	for _, book := range books {
		bookInfoResp = append(bookInfoResp, &pb.ListBooksResponse_BookInfo{
			Isbn:    book.ISBN,
			Name:    book.Name,
			Author:  book.Author,
			AddedBy: book.AddedBy,
		})
	}
	return &pb.ListBooksResponse{Books: bookInfoResp}, nil
}

func (b *BookShelfServer) DelBook(ctx context.Context, req *pb.DelBookRequest) (*pb.DelBookResponse, error) {
	r, err := b.B.Del(req.Isbn)
	if err != nil {
		return &pb.DelBookResponse{Message: "Error: "}, err
	}
	return &pb.DelBookResponse{Message: r}, nil
}

func (b *BookShelfServer) FindBook(ctx context.Context, req *pb.FindBookRequest) (*pb.FindBookResponse, error) {
	book, err := b.B.Find(req.Isbn)
	if err != nil {
		return &pb.FindBookResponse{}, err
	}
	return &pb.FindBookResponse{Book: &pb.FindBookResponse_Book{
		Isbn:    book.ISBN,
		Name:    book.Name,
		Author:  book.Author,
		AddedBy: book.AddedBy,
	}}, nil
}
