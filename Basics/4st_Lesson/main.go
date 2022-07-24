package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `jsom:"name"`
	Age  uint16 `jsom:"age"`
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	/*Усановка данных для базы
	insert, err := db.Query("INSERT INTO `uers` (`name`, `age`) VALUES('Bob', 35)")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
	*/

	//Выборка данных
	res, err := db.Query("SELECT `name`, `age` FROM `uers`")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))

	}
}
