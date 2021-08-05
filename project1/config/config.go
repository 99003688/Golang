package config

import (
	"database/sql"
	// "fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (db *sql.DB, err error) {
	// dbDriver := "mysql"
	// dbName := "db6"
	// dbUser := "root"
	// dbPass := "root"
	// fmt.Println(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/db6")

	if err != nil {
		panic(err.Error())
	}

	statement, _ :=
		db.Prepare("CREATE TABLE IF NOT EXISTS product (id INTEGER PRIMARY KEY, name TEXT, price FLOAT,quantity INTEGER,description TEXT)")
	statement.Exec()
	return

}
