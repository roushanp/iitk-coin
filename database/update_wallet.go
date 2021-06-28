package database

import (
	//"database/sql"
	//"log"
	"fmt"
	"strconv"
	//"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func AddCoin(claim_roll int, rollno int, coin int) {

	admin :=0
	DB.QueryRow("SELECT IsAdmin FROM User WHERE rollno=?", claim_roll).Scan(&admin)
	if(admin!=2){
		fmt.Println("Not allowed to award coin")
		return
	}
	mutex.Lock()
	_, err := DB.Exec("UPDATE User SET coin = coin + ? WHERE rollno=?", coin, rollno)
	checkErr(err)
	var detail string = strconv.Itoa(coin)+" coin is awarded to " + strconv.Itoa(rollno)
	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)
	_, err = DB.Exec("INSERT INTO History(time, rollno, details) VALUES (?,?,?)", dt, rollno, detail)
	checkErr(err)
	mutex.Unlock()

}

func Transfer(rollno1 int, rollno2 int, coin int) {

	mutex.Lock()
	tx, err := DB.Begin()
	checkErr(err)
	batch1 := ""
	batch2 := ""
	tax:=0
	admin := 0
	events := 0
	var coin1 int
	tx.QueryRow("SELECT batch,IsAdmin,N_events,coin FROM User WHERE rollno=?", rollno1).Scan(&batch1, &admin, &events, &coin1)
	tx.QueryRow("SELECT batch FROM User WHERE rollno=?", rollno2).Scan(&batch2)
	if((admin==0)&&(events<=0)){
		fmt.Println("You have not participated in sufficient number of events")
		tx.Rollback()
		mutex.Unlock()
		return
	}
	if((admin==0)&&(batch1==batch2)){tax=int(0.02*float64(coin))}
	if((admin==0)&&(batch1!=batch2)){tax=int(0.33*float64(coin))}
	if (coin1 - (coin+tax)) < 0 {
		fmt.Println("doing rollback1")
		tx.Rollback()
		mutex.Unlock()
		return
	}

	_, err = tx.Exec("UPDATE User SET coin = coin - ? WHERE rollno=?", coin+tax, rollno1)

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
		var detail string = strconv.Itoa(coin)+" coin is transferred from " + strconv.Itoa(rollno1)
		var datetime = time.Now()
		dt := datetime.Format(time.RFC3339)
		_, err = DB.Exec("INSERT INTO History(time, rollno, details) VALUES (?,?,?)", dt, rollno1, detail)
		checkErr(err)
		detail = strconv.Itoa(coin)+" coin is transferred to " + strconv.Itoa(rollno2)
		_, err = DB.Exec("INSERT INTO History(time, rollno, details) VALUES (?,?,?)", dt, rollno2, detail)
		checkErr(err)
		mutex.Unlock()
		return
	}
}

func Balance(rollno int) int {
	var coin int
	DB.QueryRow("SELECT coin FROM User WHERE rollno=?", rollno).Scan(&coin)
	return coin
}
