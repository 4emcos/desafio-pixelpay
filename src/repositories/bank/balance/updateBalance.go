package balance

import "database/sql"

func UpdateBalance(value float64, id string, db *sql.DB) error {
	_, err := db.Exec("UPDATE person SET balance = $1 WHERE id = $2", value, id)
	if err != nil {
		return err
	}

	return nil

}
