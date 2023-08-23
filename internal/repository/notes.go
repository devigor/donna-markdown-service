package repository

import (
	"context"
	"log"

	database "github.com/devigor/donna-markdown-service/internal/config/db"
	"github.com/devigor/donna-markdown-service/internal/contracts"
	"github.com/google/uuid"
)

func Create(content string) error {
	data := &contracts.NotesBody{
		Id:      uuid.New().String(),
		Content: content,
	}

	db, err := database.OpenConn()
	if err != nil {
		log.Fatalln("Error to connect database\n%r", err)
		return err
	}

	_, error := db.Exec(context.Background(),
		"INSERT INTO donna_notes (id, content) VALUES ($1, $2)", data.Id, data.Content)

	defer db.Close(context.Background())
	return error
}
