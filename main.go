package main

import (
	"16_DB/DB"
	"16_DB/moduls"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)



func main() {
	db, err := sql.Open("sqlite3", "Bank")
	if err != nil{
		log.Fatal("DB is not connected", err)
	}

	err = DB.DbInit(db)
	if err != nil {
		log.Fatal("error 1 ", err)
	}
	err = DB.DbInitCurrency(db)
	if err != nil {
		log.Fatal("error 2 ", err)
	}
	err = DB.DbInitAccaunts(db)
	if err != nil {
		log.Fatal("error 3", err)
	}
	//err = DB.DBInsert(db)
	//if err != nil {
	//	log.Fatal("error 4 ", err)
	//}
	fmt.Println("ok")
	//defer  db.Close()


	user := moduls.User{
		Surname:  "Nurova",
		Age:      17,
		Sex:      "female",
	}
	err = DB.AddNewClient(db, user)
	if err!= nil {
		fmt.Println("error ", err)
	}

	err = DB.UpdateUsers(db)
		if err != nil{
			fmt.Println("error in updating, bro", err)
		}

		accaunt := moduls.Accaunts{
			UserId:   3,
			Number:   22889,
			Amount:   999,
			Currency: 1,
		}
	err = DB.AddNewAccaunt(db, accaunt)
	if err != nil{
		fmt.Println(`error while adding accaunt`, err)
	}
}

//все 3 таблицы реализовать
// select, insert, update, delete, drop - read
