package balanceUseCase

import (
	"database/sql"
	"desafio-pixelpay/src/repositories/bank/balance"
	"desafio-pixelpay/src/repositories/user/exists"
)

func AccountBalance(id string, db *sql.DB) (float64, error) {
	userExists, err := exists.ExistsById(id, db)
	if err != nil {
		return 500, err
	}

	if !userExists {
		return 404, nil
	}

	accBalance, err := balance.CheckBalance(id, db)

	if err != nil {
		return 500, err

	}

	return accBalance, nil
}
