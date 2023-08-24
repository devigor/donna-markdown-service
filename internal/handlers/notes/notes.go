package notes

import (
	"context"
	"fmt"

	notes "github.com/devigor/donna-markdown-service/internal/contracts"
	"github.com/devigor/donna-markdown-service/internal/repository"
)

type notesServiceServer struct {
	notes.UnimplementedNotesServiceServer
}

func NewServer() *notesServiceServer {
	return &notesServiceServer{}
}

func (s *notesServiceServer) GetAll(ctx context.Context, request *notes.Empty) (*notes.GetAllResponse, error) {
	data, error := repository.Select()
	fmt.Println(&notes.GetAllResponse{Items: data})
	return &notes.GetAllResponse{Items: data}, error
}

func (s *notesServiceServer) CreateNote(ctx context.Context, request *notes.CreateNoteRequest) (*notes.Empty, error) {
	// create note
	error := repository.Create(request.Content)

	return &notes.Empty{}, error
}

func (s *notesServiceServer) DeleteNote(ctx context.Context, request *notes.DeleteNoteRequest) (*notes.Empty, error) {
	error := repository.Update(request.Id, request.Content)

	return &notes.Empty{}, error
}
