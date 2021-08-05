package main

import (
	"context"
	p "project5/protobuffer/hellopb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	p.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *p.Request) (*p.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &p.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *p.Request) (*p.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &p.Response{Result: result}, nil
}