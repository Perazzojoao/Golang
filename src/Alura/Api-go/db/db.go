package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

)

func ConectarDb() *sql.DB {
	db, err := sql.Open("mysql", "root:joao1998J@tcp(localhost:3306)/alura_loja")

	if err != nil {
		fmt.Println("ERROR -> Fail to validate sql.Open() arguments")
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("ERROR -> Fail to verify connection with db.Ping()")
		panic(err.Error())
	}
	return db
}
