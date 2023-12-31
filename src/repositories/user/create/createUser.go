package create

import (
	"database/sql"
	"desafio-pixelpay/src/model"
	"fmt"
	"github.com/google/uuid"
)

func CreateUser(modelUser model.User, db *sql.DB) error {
	var id = uuid.New().String()

	_, err := db.Exec("INSERT INTO person (id, name, document, email) VALUES ($1, $2, $3, $4)", &id, modelUser.Name, modelUser.Document, modelUser.Email)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
