# BookShelf Example gRPC service 

[![BlogPost](https://img.shields.io/badge/BlogPost-mrturkmen.com-brightgreen)](https://mrturkmen.com/gRPC-calls-with-evans/)

This repo is created for demonstrating basic usage of [__evans__](https://github.com/ktr0731/evans) on the blog post. 

https://mrturkmen.com/gRPC-calls-with-evans/

- [Proto File](#proto-file)
- [Compiling Proto](#compiling-proto)
- [Running Service](#running-service)
- [Demo](#demo)

## Proto File 

```proto
// it is important to declare syntax version
syntax = "proto3";
service BookShelf {
    rpc AddBook(AddBookRequest) returns (AddBookResponse) {}
    rpc ListBook (ListBooksRequest) returns (ListBooksResponse) {}
    rpc DelBook (DelBookRequest) returns (DelBookResponse){}
    rpc FindBook (FindBookRequest) returns (FindBookResponse){}
}


message AddBookRequest {
    string addedBy = 1;
    BookInfo book = 2;
    message BookInfo {
        string isbn =1;
        string name =2;
        string author=3;
        string addedBy=4;
    }

}

message AddBookResponse {
    string message = 1;
}
message ListBooksRequest {
// no need to have anything
// could be extended to list books based on category ...
}

message ListBooksResponse {
    repeated BookInfo books =1;
    message BookInfo {
        string isbn =1;
        string name =2;
        string author=3;
        string addedBy=4;
    }
}

message  DelBookRequest {
    string isbn =1;
}

message DelBookResponse {
    string message =1;
}
message FindBookRequest {
    string isbn =1;
}

message FindBookResponse {
    Book book = 1;
    message Book {
        string isbn =1;
        string name =2;
        string author=3;
        string addedBy=4;
    }
}

```

## Compiling proto 

```bash 
$ protoc -I proto/ proto/bs.proto --go_out=plugins=grpc:proto 
```

## Running server 
```bash 
$ go run server/main.go 
```

## Demo

[![Evans Demo](http://img.youtube.com/vi/GnAUkPUXYCs/0.jpg)](http://www.youtube.com/watch?v=GnAUkPUXYCs "BookShelf evans demo")
