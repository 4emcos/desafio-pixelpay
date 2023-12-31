package withdrawUseCase

import (
	"database/sql"
	"desafio-pixelpay/src/model"
	"desafio-pixelpay/src/repositories/bank/balance"
)

func WithdrawValue(withdrawModel model.Withdraw, db *sql.DB) (float64, int, error) {
	accBalance, err := balance.CheckBalance(withdrawModel.Payee, db)

	if err != nil {
		return 0, 500, err

	}
	if accBalance < withdrawModel.Value {
		return 0, 422, nil
	}

	accBalance = accBalance - withdrawModel.Value

	err = balance.UpdateBalance(accBalance, withdrawModel.Payee, db)

	if err != nil {
		return 0, 500, err
	}

	return accBalance, 200, nil

}
