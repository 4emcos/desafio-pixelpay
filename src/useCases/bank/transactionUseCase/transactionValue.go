package transactionUseCase

import (
	"database/sql"
	"desafio-pixelpay/src/model"
	"desafio-pixelpay/src/repositories/bank/balance"
	"desafio-pixelpay/src/repositories/user/exists"
)

func TransactionValue(transactionModel model.Transaction, db *sql.DB) (int, error, string) {

	peyerExists := checkUsersExists(transactionModel.Payer, transactionModel.Payee, db)

	if !peyerExists {
		return 422, nil, "payer or payee not exists"
	}

	balanceValue, err := balance.CheckBalance(transactionModel.Payer, db)

	if err != nil {
		return 500, err, ""
	}

	if balanceValue < transactionModel.Value {
		return 422, nil, "payer not have balance"
	}

	balanceValue = balanceValue - transactionModel.Value

	err = balance.UpdateBalance(balanceValue, transactionModel.Payer, db)

	if err != nil {
		return 500, err, ""
	}

	balanceValue, err = balance.CheckBalance(transactionModel.Payee, db)

	if err != nil {
		return 500, err, ""
	}

	balanceValue = balanceValue + transactionModel.Value

	err = balance.UpdateBalance(balanceValue, transactionModel.Payee, db)

	return 200, nil, ""
}

func checkUsersExists(payer string, payee string, db *sql.DB) bool {
	_, err := exists.ExistsById(payer, db)
	if err != nil {
		return false
	}

	_, err = exists.ExistsById(payee, db)
	if err != nil {
		return false
	}

	return true
}
