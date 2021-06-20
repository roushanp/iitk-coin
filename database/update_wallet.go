package database

import (
	"database/sql"
	//"log"
	"fmt"
	//"strconv"
	//"sync"
	//"time"

	_ "github.com/mattn/go-sqlite3"
)

func AddCoin(rollno int, coin int){
	db, err := sql.Open("sqlite3","./coin.db")
	checkErr(err)
	mutex.Lock()
	//time.Sleep(10000*time.Millisecond)
	stmt, err := db.Prepare("UPDATE User SET coin=coin+? WHERE rollno=?")
    checkErr(err)

    _,err = stmt.Exec(coin, rollno)

	checkErr(err)

	mutex.Unlock()
	db.Close()
}

func Transfer(rollno1 int, rollno2 int, coin int){
	db, err := sql.Open("sqlite3","./coin.db")
	checkErr(err)
	
	tx, err := db.Begin()
	checkErr(err)
	mutex.Lock()

	var coin1 int;
	db.QueryRow("SELECT coin FROM User WHERE rollno=?", rollno1).Scan(&coin1)
	if((coin1-coin)<0){
		fmt.Println("doing rollback")
		tx.Rollback()
		mutex.Unlock()
	}

	_, err = tx.Exec("UPDATE User SET coin = coin - ? WHERE rollno=? AND coin - ? >= 0", coin, rollno1, coin)

	if err != nil {
		fmt.Println("doing rollback")
		tx.Rollback()
		mutex.Unlock()
	}

	//time.Sleep(10000*time.Millisecond)
	_, err = tx.Exec("UPDATE User SET coin = coin + ? WHERE rollno=?", coin, rollno2)

	if err != nil {
        fmt.Println("doing rollback")
        tx.Rollback()
		mutex.Unlock()
    } else {
        tx.Commit()
		mutex.Unlock()
    }	

}

func Balance(rollno int) int{
	db, err := sql.Open("sqlite3","./coin.db")
	checkErr(err)
	mutex.Lock()
	var coin int
	db.QueryRow("SELECT coin FROM User WHERE rollno=?", rollno).Scan(&coin)
	mutex.Unlock()
	return coin
}
