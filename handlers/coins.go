package handlers

import(
	"fmt"
	//"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
	//"golang.org/x/crypto/bcrypt"
	//"github.com/dgrijalva/jwt-go"

	"github.com/roushanp/iitk-coin/database"
)

func Award(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	
	keyVal1 := make(map[string]int)
	keyVal2 := make(map[string]int)
	
    json.Unmarshal(body, &keyVal1) // check for errors
	json.Unmarshal(body, &keyVal2) // check for errors
	
	rollno := keyVal1["rollno"]
	coin := keyVal2["coin"]

    if(r.Method == "POST"){
		//fmt.Fprintf(w, "POST method passed in signup %d %s %s",rollno,hash,name)
		database.AddCoin(Claim_roll,rollno, coin)
	}
}

func Transfer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	
	keyVal1 := make(map[string]int)
	keyVal2 := make(map[string]int)
	keyVal3 := make(map[string]int)
	
    json.Unmarshal(body, &keyVal1) // check for errors
	json.Unmarshal(body, &keyVal2) // check for errors
	json.Unmarshal(body, &keyVal3) // check for errors
	
	rollno1 := keyVal1["rollno1"]
	rollno2 := keyVal2["rollno2"]
	coin := keyVal3["coin"]

	if(Claim_roll!=rollno1){
		fmt.Fprintf(w, "%d is not allowed to transfer coin from %d",Claim_roll,rollno1)
		return
	}

    if(r.Method == "POST"){
		//fmt.Fprintf(w, "POST method passed in signup %d %s %s",rollno,hash,name)
		database.Transfer(rollno1, rollno2, coin)
	}
}

func Balance(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	keyVal := make(map[string]int)
	json.Unmarshal(body, &keyVal) // check for errors
	rollno := keyVal["rollno"]
	coin := 0

	if(Claim_roll!=rollno){
		fmt.Fprintf(w, "%d is not allowed to get balance details of %d",Claim_roll,rollno)
		return
	}

	if(r.Method == "GET"){
		coin = database.Balance(rollno)
		fmt.Fprintf(w, "Available Coin of %d is %d",rollno,coin)
	}
}