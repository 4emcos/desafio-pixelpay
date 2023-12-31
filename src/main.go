package main

import (
	"database/sql"
	"desafio-pixelpay/src/infra"
	"desafio-pixelpay/src/routers"
	"fmt"
)

func main() {

	db, err := infra.ConnectDB()
	router := routers.InitRoutes(db)

	if err != nil {
		fmt.Println(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	port := ":8080"

	router.Run(port)
}
