package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/devigor/donna-markdown-service/contracts"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type server struct {
	contracts.UnimplementedMarkdownServer
}

func (s *server) GetAll(ctx context.Context, request *contracts.Empty) (*contracts.GetAllResponse, error) {
	data := []*contracts.MarkdownBody{
		{
			Id:        uuid.New().String(),
			Content:   "# Title",
			CreatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
			UpdatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
		},
		{
			Id:        uuid.New().String(),
			Content:   "# Title 2",
			CreatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
			UpdatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
		},
		{
			Id:        uuid.New().String(),
			Content:   "# Title 3",
			CreatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
			UpdatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
		},
	}
	return &contracts.GetAllResponse{
		Items: data,
	}, nil
}

var port = flag.Int("port", 50051, "The server port")

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v: ", err)
	}

	s := grpc.NewServer()
	contracts.RegisterMarkdownServer(s, &server{})
	log.Printf("server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
