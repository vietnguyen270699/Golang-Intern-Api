package connect

import (
	"database/sql"

	_ "github.com/godror/godror"
)

func ConnectDB() *sql.DB {

	db, err := sql.Open("godror", `user="SYSTEM" password="dev" connectString="localhost:1521/xe"`)
	if err != nil {
		panic(err)
	}
	return db

}
