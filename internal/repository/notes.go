package repository

import (
	"context"
	"log"
	"time"

	database "github.com/devigor/donna-notes-service/internal/config/db"
	"github.com/devigor/donna-notes-service/internal/contracts"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Select() ([]*contracts.NotesBody, error) {
	db, err := database.OpenConn()
	if err != nil {
		log.Fatalln("Error to connect database\n%r", err)
	}

	rows, error := db.Query(context.Background(),
		"SELECT id, content, created_at, updated_at FROM donna_notes")
	defer db.Close(context.Background())
	defer rows.Close()

	if error != nil {
		log.Fatalln(error)
	}

	var results []*contracts.NotesBody

	for rows.Next() {
		var noteStruct contracts.NotesBody
		var createdAt, updatedAt time.Time
		if err := rows.Scan(&noteStruct.Id, &noteStruct.Content, &createdAt, &updatedAt); err != nil {
			log.Fatalln(err)
		}

		noteStruct.CreatedAt = timestamppb.New(createdAt) // Converta para timestamppb.Timestamp
		noteStruct.UpdatedAt = timestamppb.New(updatedAt)

		results = append(results, &noteStruct)
	}

	return results, nil
}

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

func Update(id string, content string) error {
	db, err := database.OpenConn()
	if err != nil {
		log.Fatalln("Error to connect database\n%r", err)
	}
	defer db.Close(context.Background())

	row, error := db.Exec(context.Background(),
		"UPDATE donna_notes AS dn SET content = $1 WHERE dn.id = $2", content, id)

	if row.RowsAffected() == 0 {
		log.Fatalln("Error to update the value")
	}

	return error
}

func Delete(id string) error {
	db, err := database.OpenConn()
	if err != nil {
		log.Fatalln("Error to connect database\n%r", err)
	}
	defer db.Close(context.Background())

	row, error := db.Exec(context.Background(),
		"DELETE FROM donna_notes AS dn WHERE dn.id = $1", id)

	if row.RowsAffected() == 0 {
		log.Fatalln("Error to delete the value", error)
	}

	return error
}

func FindById(id string) (*contracts.NotesBody, error) {
	db, err := database.OpenConn()
	if err != nil {
		log.Fatalln("Error to connect database\n%r", err)
	}
	defer db.Close(context.Background())

	rows, error := db.Query(context.Background(),
		"SELECT id, content, created_at, updated_at FROM donna_notes dn WHERE dn.id = $1", id)

	defer rows.Close()

	var results *contracts.NotesBody

	for rows.Next() {
		var noteStruct contracts.NotesBody
		var createdAt, updatedAt time.Time
		if err := rows.Scan(&noteStruct.Id, &noteStruct.Content, &createdAt, &updatedAt); err != nil {
			log.Fatalln(err)
		}

		noteStruct.CreatedAt = timestamppb.New(createdAt) // Converta para timestamppb.Timestamp
		noteStruct.UpdatedAt = timestamppb.New(updatedAt)

		results = &noteStruct
	}
	return results, error
}
