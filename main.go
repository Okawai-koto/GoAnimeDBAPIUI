package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	db, err := GetDB()
	if err != nil {
		fmt.Println("err1", err)

	} else {
		response, err := db.Query("CREATE TABLE Persons23 (PersonID int,LastName varchar(255),FirstName varchar(255),Address varchar(255),City varchar(255));")
		if err != nil {
			fmt.Println("err2", err)
		} else {
			fmt.Println("%s", response)
		}
	}
}

func GetDB() (db *sql.DB, err error) {
	db, err = sql.Open("mssql", "server=localhost; Trusted_Connection=true;Database=GoAnimeDB")
	return
}
