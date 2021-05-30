package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)
type person struct{
	rollno int
	name string
}

func addUser(a int, s string){
	db, _ := sql.Open("sqlite3","./coin.db")
	statement, _ := db.Prepare("INSERT INTO User (rollno, name) VALUES (?, ?)")
	statement.Exec(a,s)
}

func main() {
	db, _ := sql.Open("sqlite3","./coin.db")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS User (rollno INTEGER PRIMARY KEY, name TEXT)")
	statement.Exec()
	var user = []person { 
		 {
			rollno: 190721, 
			name: "Roushan",
		},
		 {
			rollno: 190722, 
			name: "Roushan2",
		},{
			rollno: 190723, 
			name: "Roushan3",
		},
	}
	for i := 0; i < len(user); i++{
		addUser(user[i].rollno,user[i].name)
	}
	
	/*rows, _ := db.Query("SELECT rollno, name FROM User")
    var rollno int
    var name string
    for rows.Next() {
        rows.Scan(&rollno, &name)
        fmt.Println(strconv.Itoa(rollno) + ": " + name)
    }*/
	
}