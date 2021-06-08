package database

import (
	"database/sql"
	"log"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)
func checkErr(err error){
	if(err!=nil){log.Fatal(err)}
}

func addUser(a int, s string){
	db, err := sql.Open("sqlite3","./coin.db")
	checkErr(err)
	statement, err := db.Prepare("INSERT INTO User (rollno, password) VALUES (?, ?)")
	checkErr(err)
	statement.Exec(a,s)
}

func GetPassword(rollno int) string{
	db, err := sql.Open("sqlite3","./coin.db")
	checkErr(err)
	var password string
	db.QueryRow("SELECT password FROM User WHERE rollno=?", rollno).Scan(&password)
	return password
}

func Insert(rollno int, password string) {
	db, err := sql.Open("sqlite3","./coin.db")
	checkErr(err)
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS User (rollno INTEGER PRIMARY KEY, password TEXT)")
	checkErr(err)
	statement.Exec()
	addUser(rollno, password)
	
	rows, err := db.Query("SELECT rollno, password FROM User")
	checkErr(err)
    for rows.Next() {
        rows.Scan(&rollno, &password)
        fmt.Println(strconv.Itoa(rollno) + ": " + password)
    }
	db.Close()
	
}