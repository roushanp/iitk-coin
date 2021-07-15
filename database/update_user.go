package database

import (
	"database/sql"
	"log"
	//"fmt"
	//"strconv"
	"sync"
	//"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

var mutex = &sync.Mutex{}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() {
	db, err := sql.Open("sqlite3", "./coin.db")
	checkErr((err))

	DB = db

	statement, err := DB.Prepare("CREATE TABLE IF NOT EXISTS User (rollno INTEGER PRIMARY KEY, name TEXT, batch TEXT, IsAdmin INTEGER, N_events INTEGER, coin INTEGER)")
	checkErr(err)
	statement.Exec()
	statement, err = DB.Prepare("CREATE TABLE IF NOT EXISTS Auth (rollno INTEGER PRIMARY KEY, password TEXT)")
	checkErr(err)
	statement.Exec()
	statement, err = DB.Prepare("CREATE TABLE IF NOT EXISTS History (time DATETIME DEFAULT CURRENT_TIMESTAMP, rollno INTEGER, details TEXT)")
	checkErr(err)
	statement.Exec()
	statement, err = DB.Prepare("CREATE TABLE IF NOT EXISTS RedeemPending (time DATETIME DEFAULT CURRENT_TIMESTAMP, rollno INTEGER, item_name TEXT, coin INTEGER)")
	checkErr(err)
	statement.Exec()
	statement, err = DB.Prepare("CREATE TABLE IF NOT EXISTS RedeemItem (itemName TEXT PRIMARY KEY,itemLeft INTEGER, itemCost INTEGER)")
	checkErr(err)
	statement.Exec()
	statement, err = DB.Prepare("DROP TABLE IF EXISTS ReedemItem")
	checkErr(err)
	statement.Exec()
	/*coin := 0

		rows, err := DB.Query("SELECT * FROM User")
		checkErr(err)
	    for rows.Next() {
	        rows.Scan(&rollno, &name, &batch, &IsAdmin, &coin)
	        fmt.Println(strconv.Itoa(rollno) + ": " + name+ " "+ batch+ " "+strconv.Itoa(IsAdmin)+" "+ strconv.Itoa(coin))
	    }
		rows, err = DB.Query("SELECT * FROM Auth")
		checkErr(err)
	    for rows.Next() {
	        rows.Scan(&rollno, &password)
	        fmt.Println(strconv.Itoa(rollno) + ": " + password)
	    }*/
}

func AddUser(rollno int, name string, batch string, IsAdmin int, N_events int, password string) {
	coin := 0
	statement, err := DB.Prepare("INSERT INTO User (rollno, name, batch, IsAdmin, N_events, coin) VALUES (?, ?, ?, ?, ?, ?)")
	checkErr(err)
	statement.Exec(rollno, name, batch, IsAdmin, N_events, coin)
	statement, err = DB.Prepare("INSERT INTO Auth (rollno, password) VALUES (?, ?)")
	checkErr(err)
	statement.Exec(rollno, password)
}

func GetPassword(rollno int) string {
	var password string
	DB.QueryRow("SELECT password FROM Auth WHERE rollno=?", rollno).Scan(&password)
	return password
}

func GetUserDetails(rollno int) (string, string) {
	var name string
	var batch string
	DB.QueryRow("SELECT name,batch FROM User WHERE rollno=?", rollno).Scan(&name, &batch)
	return name, batch
}
