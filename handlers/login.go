package handlers

import(
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/roushanp/iitk-coin/database"
)

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
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
		w.Header().Set("Content-Type", "application/json") 
      	user := User {
                	Rollno: rollno, 
                    Jwtoken: tokenString,
				}
      
     	json.NewEncoder(w).Encode(user) 
		//fmt.Fprintf(w, "Password matched %s",tokenString)
	}
	
}