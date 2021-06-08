package main

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/roushanp/iitk-coin/database"
)
var jwtKey = []byte("my_secret_key")

type Claims struct {
	Roll int `json:"rollno"`
	jwt.StandardClaims
}

func checkErr(err error){
	if(err!=nil){log.Fatal(err)}
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
	return string(hash)
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	/*if r.URL.Path != "/signup" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }*/
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	
	keyVal1 := make(map[string]int)
    keyVal2 := make(map[string]string)
	
    json.Unmarshal(body, &keyVal1) // check for errors
	json.Unmarshal(body, &keyVal2) // check for errors
	
	rollno := keyVal1["rollno"]
    password := keyVal2["password"]
	hash := hashAndSalt([]byte(password))

    if(r.Method == "POST"){
		fmt.Fprintf(w, "POST method passed in signup %d %s",rollno,hash)
		database.Insert(rollno, hash)
	}
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
	err = bcrypt.CompareHashAndPassword(password_db,password)
	if err != nil {
        log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
    }
	if(err==nil){
		expirationTime := time.Now().Add(5000 * time.Minute)
		claims := &Claims{
			Roll: rollno,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
		fmt.Fprintf(w, "Password matched %s",tokenString)
	}
	
}

func SecretPage(w http.ResponseWriter, r *http.Request){
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
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
	fmt.Fprintf(w, "Welcome to IITK")
}

func main() {

	http.HandleFunc("/signup", SignUpHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/secretpage", SecretPage)
	
    fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}