package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	contracts "github.com/devigor/donna-markdown-service/internal/protos"
	"google.golang.org/grpc"
)

type server struct {
	contracts.UnimplementedNotesServer
}

var port = flag.Int("port", 50051, "The server port")

func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v: ", err)
	}

	s := grpc.NewServer()
	contracts.RegisterNotesServer(s, &server{})

	log.Printf("server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
