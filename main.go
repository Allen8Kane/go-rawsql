package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	id       int
	username string
	salary   int
}

func main() {
	database, _ := sql.Open("sqlite3", "./dev.db")
	sql_create := `CREATE TABLE Users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		salary INTEGER NOT NULL
	)`
	statement, err := database.Prepare(sql_create)
	if err != nil {
		panic(err)
	}
	statement.Exec()
	sql_insert := `
		INSERT INTO Users (username,salary) 
		VALUES(?, ?)
	`
	statement, _ = database.Prepare(sql_insert)
	statement.Exec("Alex", 350)
	sql_select := `
		SELECT * FROM Users
	`
	rows, _ := database.Query(sql_select)
	users := []User{}
	for rows.Next() {
		user := User{}
		rows.Scan(&user.id, &user.username, &user.salary)
		users = append(users, user)
	}
	for _, user := range users {
		fmt.Println(user.id, user.username, user.salary)
	}
}
