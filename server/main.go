package main

import (
	"fmt"

	"github.com/mrturkmencom/bookshelf/store"

	"log"
	"net"

	pb "github.com/mrturkmencom/bookshelf/proto"
	b "github.com/mrturkmencom/bookshelf/store"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	bookStore := store.NewInMemoryStore()
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	bookShelfServer := b.BookShelfServer{B: bookStore}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterBookShelfServer(grpcServer, &bookShelfServer)
	fmt.Println("BookShelf gRPC server is running ....")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
