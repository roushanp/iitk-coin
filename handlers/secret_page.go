package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"

	"github.com/roushanp/iitk-coin/database"
)

var Claim_roll int = 0
var tkn *jwt.Token

func CheckToken(){
	if !tkn.Valid {
		fmt.Println("Token Expired")
		return
	}
}

func SecretPage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err = jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	Claim_roll = claims.Roll
	name, batch := database.GetUserDetails(claims.Roll)
	fmt.Fprintf(w, "Welcome to IITK Coin %s. Your batch is %s and you have succesfully logged in our system. Now you can access award, transfer, and balance endpoints", name, batch)
	http.HandleFunc("/award", Award)
	http.HandleFunc("/transfer", Transfer)
	http.HandleFunc("/balance", Balance)
	http.HandleFunc("/redeem",Redeem)
	http.HandleFunc("/redeemProc",RedeemProc)
	http.HandleFunc("/additem",AddItem)
}
