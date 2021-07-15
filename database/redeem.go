package database

import (
	//"database/sql"
	//"log"
	"fmt"
	"strconv"
	//"sync"
	"time"
	//_ "github.com/mattn/go-sqlite3"
)

func AddItem(claim_roll int, itemName string, itemLeft int, itemCost int){
	mutex.Lock()
	tx, err := DB.Begin()
	checkErr(err)
	admin1 := -1
	tx.QueryRow("SELECT IsAdmin FROM User WHERE rollno=?", claim_roll).Scan(&admin1)
	if admin1 != 2 {
		fmt.Println("Only GenSec and Associate Heads are allowed to add items for Redeem")
		tx.Rollback()
		mutex.Unlock()
		return
	}
	stmt, err := tx.Prepare("INSERT INTO RedeemItem (itemName, itemLeft, itemCost) VALUES (?, ?, ?)")
	checkErr(err)
	stmt.Exec(itemName, itemLeft, itemCost)
	tx.Commit()
	mutex.Unlock()
}

func RedeemReq(claim_roll int, coin int, item string) {
	mutex.Lock()
	tx, err := DB.Begin()
	checkErr(err)
	admin1 := -1
	tx.QueryRow("SELECT IsAdmin FROM User WHERE rollno=?", claim_roll).Scan(&admin1)
	if admin1 != 0 {
		fmt.Println("Only students are allowed to put redeem requests")
		tx.Rollback()
		mutex.Unlock()
		return
	}
	itemleft:=0
	tx.QueryRow("SELECT itemLeft FROM RedeemItem WHERE itemName=?", item).Scan(&itemleft)
	if(itemleft==0){
		fmt.Println("No more items are left")
		tx.Rollback()
		mutex.Unlock()
		return
	}

	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)
	_, err = tx.Exec("INSERT INTO RedeemPending(time, rollno, item_name, coin) VALUES (?,?,?,?)", dt, claim_roll, item, coin)
	checkErr(err)
	if err != nil {
		fmt.Println("Doing rollback1")
		tx.Rollback()
		mutex.Unlock()
		return
	}
	_, err = tx.Exec("UPDATE RedeemItem SET itemLeft = itemLeft - 1 WHERE itemName=?", item)
	if err != nil {
		fmt.Println("Doing rollback2")
		tx.Rollback()
		mutex.Unlock()
		return
	}
	tx.Commit()
	mutex.Unlock()
}
func RedeemProc(claim_roll int, rollno int, item string, accept int) {
	mutex.Lock()
	tx, err := DB.Begin()
	checkErr(err)
	admin := -1
	tx.QueryRow("SELECT IsAdmin FROM User WHERE rollno=?", claim_roll).Scan(&admin)
	if admin != 2 {
		fmt.Println("Only GenSec and Associate Heads are allowed to process redeem requests")
		tx.Rollback()
		mutex.Unlock()
		return
	}
	cost := 0
	tx.QueryRow("SELECT itemCost FROM RedeemItem WHERE itemName=?", item).Scan(&cost)
	coin := 0
	tx.QueryRow("SELECT coin FROM User WHERE rollno=?", rollno).Scan(&coin)
	if coin < cost {
		accept = 0
		fmt.Println("Enough coin is not available to redeem")
	}

	if accept == 1 {
		_, err = tx.Exec("UPDATE User SET coin = coin - ? WHERE rollno=?", cost, rollno)
		checkErr(err)
		if err != nil {
			fmt.Println("doing rollback")
			tx.Rollback()
			mutex.Unlock()
			return
		}
		var detail string = item + " having cost " + strconv.Itoa(cost) + " is redeemed"
		var datetime = time.Now()
		dt := datetime.Format(time.RFC3339)
		_, err = tx.Exec("INSERT INTO History(time, rollno, details) VALUES (?,?,?)", dt, rollno, detail)
		checkErr(err)
	} else {
		var detail string = "Redeem of " + item + " having cost " + strconv.Itoa(cost) + "failed"
		var datetime = time.Now()
		dt := datetime.Format(time.RFC3339)
		_, err = tx.Exec("INSERT INTO History(time, rollno, details) VALUES (?,?,?)", dt, rollno, detail)
		checkErr(err)
	}

	_, err = tx.Exec("DELETE FROM RedeemPending WHERE (rollno=? AND item_name=?)", rollno, item)
	checkErr(err)
	if err != nil {
		fmt.Println("doing rollback")
		tx.Rollback()
		mutex.Unlock()
		return
	}
	tx.Commit()
	mutex.Unlock()

}
