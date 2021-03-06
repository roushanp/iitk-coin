package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"

	"github.com/roushanp/iitk-coin/database"
)

var Claim_roll int = 0
var tkn *jwt.Token

func CheckToken(w http.ResponseWriter, r *http.Request)bool{
	if !tkn.Valid {
		fmt.Println("Token Expired, Please login again")
		return true
	}
	/*if (r.Cookies() == nil){
		fmt.Println("Logged out, Please login again")
		return true
	}*/

	_, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println("Logged out, Please login again")
			return true
		}
		w.WriteHeader(http.StatusBadRequest)
		return true
	}

	return false
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
	fmt.Fprintf(w, "Welcome to IITK Coin %s. Your batch is %s and you have succesfully logged in our system. Now you can access award, transfer, balance, and other endpoints", name, batch)
	
}
