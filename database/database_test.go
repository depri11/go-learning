package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseSetup(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_learn")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connect database Success")

	defer db.Close()

}