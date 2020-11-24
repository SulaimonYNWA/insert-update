package DB

import (
	"16_DB/moduls"
	"database/sql"
	"fmt"
)

func init()  {
	fmt.Println("hello, I am test")
}

func DbInit(db *sql.DB) (err error) {
	const usersDB = `CREATE TABLE if not exists user (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text not null,
    surname TEXT NOT NULL,
    age INTEGER NOT NULL,
    sex TEXT,
    remove BOOLEAN NOT NULL DEFAULT FALSE
)`
	_, err = db.Exec(usersDB)

	if err != nil {
		return err
	}
	return nil
}

func DbInitCurrency(db *sql.DB) (err error) {
	const currency = `CREATE TABLE if not exists currency (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text
)`
	_, err = db.Exec(currency)

	if err != nil {
		return err
	}
	return nil
	}

	func DbInitAccaunts(db *sql.DB) (err error) {
		const accaunts = `CREATE TABLE if not exists accaunts (
    	id integer PRIMARY KEY AUTOINCREMENT,
    	userId integer references user(id) not null ,
    	number integer,
    	amount integer,
    	currency integer references currency(id),
    	remove BOOLEAN NOT NULL DEFAULT FALSE
)`
		_, err = db.Exec(accaunts)

		if err != nil {
			return nil
		}

		return nil
	}

		func DBInsert(db *sql.DB) (err error) {
			const addToUser = `insert into clients values (3, 'NotBot', 37, 'male', false)`

			_, err = db.Exec(addToUser)

			if err != nil {
				return err
			}
			return nil

		}

			func AddNewClient(db *sql.DB, client moduls.User) (err error){
				_, err = db.Exec(`insert into user (surname, age, sex) values (($1),($2),($3))`, client.Surname, client.Age, client.Sex)
				if err != nil{
					fmt.Println("Cant insert", err)
					return err
				}
				return err
			}
			//update func for all users some columns (remove: true)
			// add new Accaunt

		func AddNewAccaunt(db *sql.DB, client moduls.Accaunts) (err error){
			_, err = db.Exec(`insert into accaunts (userId, number, amount, currency) values (($1),($2),($3),($4))`, client.UserId, client.Number, client.Amount, client.Currency)
			if err != nil{
				fmt.Println(`Cannot insert accauntt`, err)
				return err
			}
			return err
		}


func UpdateUsers(db *sql.DB) (err error){
	_, err = db.Exec(`update user set remove = false where sex = 'female' `)
	if err != nil{
		fmt.Println("cannot update the row", err)
		return err
	}
	_, err = db.Exec(`update user set remove = true where sex = 'male' & id<6 `)
	if err != nil{
		fmt.Println("cannot update the row", err)
		return err
	}
	return err
}



