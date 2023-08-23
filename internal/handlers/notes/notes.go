package notes

import (
	"context"

	notes "github.com/devigor/donna-markdown-service/internal/contracts"
	"github.com/devigor/donna-markdown-service/internal/repository"
)

type notesServiceServer struct {
	notes.UnimplementedNotesServiceServer
}

func NewServer() *notesServiceServer {
	return &notesServiceServer{}
}

func (s *notesServiceServer) CreateNote(ctx context.Context, request *notes.CreateNoteRequest) (*notes.Empty, error) {
	// create user
	error := repository.Create(request.Content)

	return &notes.Empty{}, error
}
