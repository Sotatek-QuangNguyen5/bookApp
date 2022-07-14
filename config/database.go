package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "127.0.0.1"
	port     = 3306
	user     = "root"
	password = "123qwerty"
	dbname   = "bookapp"
)

var DB *sql.DB

func getDatabase() *sql.DB {

	mysql := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)
	
	db, err := sql.Open("mysql", mysql)

	if err != nil {

		fmt.Println(err)
		return nil
	}
	return db
}

func InitDatabase() {

	DB = getDatabase()
}
