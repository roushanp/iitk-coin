package handlers

import(
	// "fmt"
	//"log"
	 "encoding/json"
	 "io/ioutil"
	 "net/http"
	//"golang.org/x/crypto/bcrypt"
	//"github.com/dgrijalva/jwt-go"

	 "github.com/roushanp/iitk-coin/database"
)

func AddItem(w http.ResponseWriter, r *http.Request){
	if(CheckToken(w,r)){return}
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	keyVal1 := make(map[string]string)
	keyVal2 := make(map[string]int)
	keyVal3 := make(map[string]int)

	json.Unmarshal(body, &keyVal1) // check for errors
	json.Unmarshal(body, &keyVal2) // check for errors
	json.Unmarshal(body, &keyVal3) // check for errors

	itemName := keyVal1["itemName"]
	itemLeft := keyVal2["itemLeft"]
	itemCost := keyVal2["itemCost"]

	if r.Method == "POST" {
		//fmt.Fprintf(w, "Redeem Request passed in Redeem %d %s",coin,item)
		database.AddItem(Claim_roll, itemName, itemLeft, itemCost)
	}
}

func Redeem(w http.ResponseWriter, r *http.Request) {
	if(CheckToken(w,r)){return}
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	keyVal1 := make(map[string]int)
	keyVal2 := make(map[string]string)

	json.Unmarshal(body, &keyVal1) // check for errors
	json.Unmarshal(body, &keyVal2) // check for errors

	coin := keyVal1["coin"]
	item := keyVal2["item"]

	if r.Method == "POST" {
		//fmt.Fprintf(w, "Redeem Request passed in Redeem %d %s",coin,item)
		database.RedeemReq(Claim_roll, coin, item)
	}
}
func RedeemProc(w http.ResponseWriter, r *http.Request){
	if(CheckToken(w,r)){return}
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	keyVal1 := make(map[string]int)
	keyVal2 := make(map[string]string)
	keyVal3 := make(map[string]int)

	json.Unmarshal(body, &keyVal1) // check for errors
	json.Unmarshal(body, &keyVal2) // check for errors
	json.Unmarshal(body, &keyVal3) // check for errors

	rollno := keyVal1["rollno"]
	item := keyVal2["item"]
	accept := keyVal3["accept"]

	if r.Method == "POST" {
		//fmt.Fprintf(w, "Redeem Process passed in Redeem %d %s",coin,item)
		database.RedeemProc(Claim_roll, rollno, item,accept)
	}
}