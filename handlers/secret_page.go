package handlers

import(
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"

	"github.com/roushanp/iitk-coin/database"
)

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