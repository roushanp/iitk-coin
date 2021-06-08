package handlers
import(
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"

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
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	
	keyVal1 := make(map[string]int)
    keyVal2 := make(map[string]string)
	keyVal3 := make(map[string]string)
	keyVal4 := make(map[string]int)
	keyVal5 := make(map[string]string)
	
    json.Unmarshal(body, &keyVal1) // check for errors
	json.Unmarshal(body, &keyVal2) // check for errors
	json.Unmarshal(body, &keyVal3) // check for errors
	json.Unmarshal(body, &keyVal4) // check for errors
	json.Unmarshal(body, &keyVal5) // check for errors
	
	rollno := keyVal1["rollno"]
	name := keyVal2["name"]
	batch := keyVal3["batch"]
	IsAdmin := keyVal4["IsAdmin"]
    password := keyVal5["password"]
	hash := hashAndSalt([]byte(password))

    if(r.Method == "POST"){
		//fmt.Fprintf(w, "POST method passed in signup %d %s %s",rollno,hash,name)
		database.Insert(rollno, name, batch, IsAdmin, hash)
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
		//fmt.Fprintf(w, "Password matched %s",tokenString)
	}
	
}

func SecretPage(w http.ResponseWriter, r *http.Request){
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
	name,batch := database.GetUserDetails(claims.Roll)
	fmt.Fprintf(w, "Welcome to IITK Coin %s. Your batch is %s and you have succesfully logged in our system",name, batch)
}

