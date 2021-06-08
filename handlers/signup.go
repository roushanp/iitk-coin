package handlers
import(
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"

	"github.com/roushanp/iitk-coin/database"
)


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
		fmt.Fprintf(w, "POST method passed in signup %d %s %s",rollno,hash,name)
		database.Insert(rollno, name, batch, IsAdmin, hash)
	}
}