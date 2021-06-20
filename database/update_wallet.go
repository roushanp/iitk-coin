package database

import (
	//"database/sql"
	//"log"
	"fmt"
	//"strconv"
	//"sync"
	//"time"

	_ "github.com/mattn/go-sqlite3"
)

func AddCoin(rollno int, coin int){

	_, err := DB.Exec("UPDATE User SET coin = coin + ? WHERE rollno=?", coin, rollno)
    checkErr(err)

}

func Transfer(rollno1 int, rollno2 int, coin int){
	
	tx, err := DB.Begin()
	checkErr(err)
	mutex.Lock()

	var coin1 int;
	tx.QueryRow("SELECT coin FROM User WHERE rollno=?", rollno1).Scan(&coin1)
	if((coin1-coin)<0){
		fmt.Println("doing rollback1")
		tx.Rollback()
		mutex.Unlock()
		return
	}

	_, err = tx.Exec("UPDATE User SET coin = coin - ? WHERE rollno=?", coin, rollno1)

	if err != nil {
		fmt.Println("doing rollback2")
		tx.Rollback()
		mutex.Unlock()
		return
	}

	//time.Sleep(10000*time.Millisecond)
	_, err = tx.Exec("UPDATE User SET coin = coin + ? WHERE rollno=?", coin, rollno2)

	if err != nil {
        fmt.Println("doing rollback3")
        tx.Rollback()
		mutex.Unlock()
		return
    } else {
        tx.Commit()
		mutex.Unlock()
		return
    }
}

func Balance(rollno int) int{
	var coin int
	DB.QueryRow("SELECT coin FROM User WHERE rollno=?", rollno).Scan(&coin)
	return coin
}
