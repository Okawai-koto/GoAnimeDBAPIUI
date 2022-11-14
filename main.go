package main

import (
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	//createTable(columnsAnime())

	tickTime := time.NewTicker(time.Second * 1)
	sayac := 0
	_ = tickTime.C
	for v := range tickTime.C {
		_ = v
		x := getApi(sayac)
		if x.Data.Id != 0 {
			fmt.Println(x.Data.Title)
			AddToDB(columnsAnime(), x)
		}
		sayac = sayac + 1
	}

	/*
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
		}*/
}
