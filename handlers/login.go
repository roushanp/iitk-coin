package handlers

import (
	"fmt"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/roushanp/iitk-coin/database"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Roll int `json:"rollno"`
	jwt.StandardClaims
}

type User struct {
	Rollno  int    `json:"rollno"`
	Jwtoken string `json:"jwt"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	keyVal1 := make(map[string]int)
	keyVal2 := make(map[string]string)

	json.Unmarshal(body, &keyVal1) // check for errors
	json.Unmarshal(body, &keyVal2) // check for errors

	rollno := keyVal1["rollno"]
	password := []byte(keyVal2["password"])
	password_db := []byte(database.GetPassword(rollno))
	err = bcrypt.CompareHashAndPassword(password_db, password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
	}
	if err == nil {
		expirationTime := time.Now().Add(10 * time.Minute)
		claims := &Claims{
			Roll: rollno,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  expirationTime,
			HttpOnly: true,
		})
		w.Header().Set("Content-Type", "application/json")
		user := User{
			Rollno:  rollno,
			Jwtoken: tokenString,
		}

		json.NewEncoder(w).Encode(user)
		//fmt.Fprintf(w, "Password matched %s",tokenString)
	}

}

func LogoutHandler(w http.ResponseWriter, r *http.Request){
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		MaxAge: -1,
	})
	w.Write([]byte("Old cookie deleted. Logged out!\n"))
	fmt.Fprintf(w, "Logout successful")
}
