package Users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"GoCasst/Auth"
	"GoCasst/Cass"

	"github.com/gocql/gocql"
)

// Put to the cassandra or log proper errors
func PostTheUser(w http.ResponseWriter, r *http.Request) {
	//var posted bool
	var errs []string
	var gocqlUUID gocql.UUID

	user, errs := ProcessUserFrom(r)

	if len(errs) == 0 && Auth.AuthCheck(w, r) {
		fmt.Println(Cass.LogTime() + ", Creating an user..")
		gocqlUUID = gocql.TimeUUID()
		// cassandra writing
		if err := Cass.Session.Query(`
    INSERT INTO users (id, name, lastname, email, birthyear, city ) VALUES (?, ?, ?, ?, ?, ?)`,
			gocqlUUID, user.Name, user.LastName, user.Email, user.BirthYear, user.City).Exec(); err != nil {
			errs = append(errs, err.Error())
		} else {
			//	posted = true
			json.NewEncoder(w).Encode(GetNewUser{ID: gocqlUUID})
			fmt.Println(Cass.LogTime()+", Created user with id: ", gocqlUUID)
		}
	} else if !Auth.AuthCheck(w, r) {
		json.NewEncoder(w).Encode(ResponseStatus{Status: "Unauthorized", Code: 401})
		fmt.Println(Cass.LogTime() + ", Post User Request - Unauthorized")
		return
	} else {
		fmt.Println(Cass.LogTime()+", Something went amiss: ", errs)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
