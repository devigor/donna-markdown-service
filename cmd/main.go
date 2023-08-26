package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	contracts "github.com/devigor/donna-notes-service/internal/contracts"
	"github.com/devigor/donna-notes-service/internal/handlers/notes"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "The server port")

func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v: ", err)
	}

	var grpcOpts []grpc.ServerOption
	server := grpc.NewServer(grpcOpts...)
	contracts.RegisterNotesServiceServer(server, notes.NewServer())

	log.Printf("server listening at %v", listen.Addr())
	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
