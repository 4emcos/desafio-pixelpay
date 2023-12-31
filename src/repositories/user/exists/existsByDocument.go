package exists

import "database/sql"

func ExistsByDocument(document string, db *sql.DB) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM person WHERE document = $1)", document).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil

}
