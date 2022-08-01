package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func getDatabase() *sql.DB {

	envMap, err := godotenv.Read("./.env")
	if err != nil {

		return nil
	}
	var (

		user = envMap["USER"]
		password = envMap["PASSWORD"]
		host = envMap["HOST"]
		port = envMap["PORTDB"]
		dbname = envMap["DBNAME"]
	)
	mysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	
	db, err := sql.Open("mysql", mysql)

	if err != nil {

		panic("Error connect to database!!!")
	}
	return db
}

func InitDatabase() {

	DB = getDatabase()
}
