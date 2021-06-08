package database

import (
	"database/sql"
	"log"
	//"fmt"
	//"strconv"

	_ "github.com/mattn/go-sqlite3"
)
func checkErr(err error){
	if(err!=nil){log.Fatal(err)}
}

func addUser(rollno int,name string, batch string, IsAdmin int, password string ){
	db, err := sql.Open("sqlite3","./coin.db")
	checkErr(err)
	statement, err := db.Prepare("INSERT INTO User (rollno, name, batch, IsAdmin) VALUES (?, ?, ?, ?)")
	checkErr(err)
	statement.Exec(rollno, name, batch, IsAdmin)
	statement, err = db.Prepare("INSERT INTO Auth (rollno, password) VALUES (?, ?)")
	checkErr(err)
	statement.Exec(rollno, password)
}

func GetPassword(rollno int) string{
	db, err := sql.Open("sqlite3","./coin.db")
	checkErr(err)
	var password string
	db.QueryRow("SELECT password FROM Auth WHERE rollno=?", rollno).Scan(&password)
	return password
}

func GetUserDetails(rollno int)(string, string){
	db, err := sql.Open("sqlite3","./coin.db")
	checkErr(err)
	var name string
	var batch string
	db.QueryRow("SELECT name,batch FROM User WHERE rollno=?", rollno).Scan(&name,&batch)
	return name,batch
}

func Insert(rollno int, name string, batch string, IsAdmin int , password string) {
	db, err := sql.Open("sqlite3","./coin.db")
	checkErr(err)
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS User (rollno INTEGER PRIMARY KEY, name TEXT, batch TEXT, IsAdmin INTEGER)")
	checkErr(err)
	statement.Exec()
	statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS Auth (rollno INTEGER PRIMARY KEY, password TEXT)")
	checkErr(err)
	statement.Exec()
	addUser(rollno, name, batch, IsAdmin, password)
	
	/*rows, err := db.Query("SELECT * FROM User")
	checkErr(err)
    for rows.Next() {
        rows.Scan(&rollno, &name, &batch, &IsAdmin)
        fmt.Println(strconv.Itoa(rollno) + ": " + name+ " "+ batch+ " "+strconv.Itoa(IsAdmin))
    }
	rows, err = db.Query("SELECT * FROM Auth")
	checkErr(err)
    for rows.Next() {
        rows.Scan(&rollno, &password)
        fmt.Println(strconv.Itoa(rollno) + ": " + password)
    }*/
	db.Close()
	
}