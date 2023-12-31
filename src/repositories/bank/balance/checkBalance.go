package balance

import "database/sql"

func CheckBalance(document string, db *sql.DB) (float64, error) {

	var balance float64

	err := db.QueryRow("SELECT balance FROM person WHERE document = $1", document).Scan(&balance)
	if err != nil {
		return 0, err

	}

	return balance, nil
}
