package createUserUseCase

import (
	"database/sql"
	"desafio-pixelpay/src/model"
	"desafio-pixelpay/src/repositories/user/create"
	"desafio-pixelpay/src/repositories/user/exists"
)

func CreateUser(userModel model.User, db *sql.DB) (int, error) {

	haveUsr, err := exists.ExistsByDocument(userModel.Document, db)

	if err != nil {
		return 500, err
	}

	if haveUsr {
		return 422, nil
	}

	err = create.CreateUser(userModel, db)

	if err != nil {
		return 500, err
	}

	return 201, nil

}
