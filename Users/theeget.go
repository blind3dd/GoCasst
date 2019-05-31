package users

import (
	"GoCasst/Cass"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blind3dd/gocasst/auth"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func GetOneUserRequest(w http.ResponseWriter, r *http.Request) {
	var exists bool = false
	var user User
	var errs []string
	vars := mux.Vars(r)
	// get all variables first for handler of concrete id
	id := vars["users_uuid"]
	uuid, err := gocql.ParseUUID(id)
	if err != nil {
		errs = append(errs, err.Error())
	} else {
		//not only pass the string of the query, but each parameter to pass into the query to fill each question mark
		m := map[string]interface{}{}
		query := "SELECT id, name, lastname, city, email, birthyear FROM users WHERE id=? LIMIT 1"
		// tell Cassandra to have single-node consistency when fetching data
		iterator := Cass.Session.Query(query, uuid).Consistency(gocql.One).Iter()
		for iterator.MapScan(m) {
			exists = true
			user = User{
				ID:        m["id"].(gocql.UUID),
				Name:      m["name"].(string),
				LastName:  m["lastname"].(string),
				Email:     m["email"].(string),
				City:      m["city"].(string),
				BirthYear: m["birthyear"].(int),
			}
		}
		if !exists || !auth.AuthCheck(w, r) {
			json.NewEncoder(w).Encode(ResponseStatus{Status: "Unauthorized", Code: 401})
			return
		} else {
			json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
			errs = append(errs, "User does not exist")
		}

		if exists {
			if auth.AuthCheck(w, r) {
				json.NewEncoder(w).Encode(GetUser{User: user})
				fmt.Println(Cass.LogTime()+", Getting user's details: ", user.ID)
				//		return
			} else {
				json.NewEncoder(w).Encode(ResponseStatus{Status: "Unauthorized", Code: 401})
				return
			}
		}
	}
}

func GetAllUsersRequest(w http.ResponseWriter, r *http.Request) {
	var listAll []User
	// some elements are not string (then interface)
	m := map[string]interface{}{}

	query := "SELECT id, name, lastname, city, email, birthyear FROM users"
	iterator := Cass.Session.Query(query).Iter()
	for iterator.MapScan(m) {
		// converting values inside Go interface - it goes with sth like int(var):
		listAll = append(listAll, User{
			ID:        m["id"].(gocql.UUID),
			Name:      m["name"].(string),
			LastName:  m["lastname"].(string),
			Email:     m["email"].(string),
			City:      m["city"].(string),
			BirthYear: m["birthyear"].(int),
		})
		m = map[string]interface{}{}
	}
	if auth.AuthCheck(w, r) {
		fmt.Println(Cass.LogTime() + ", Getting users list ..")
		json.NewEncoder(w).Encode(GetAllUsers{Users: listAll})
	} else {
		fmt.Println(Cass.LogTime() + ", Geting Users list - Unauthorized")
		json.NewEncoder(w).Encode(ResponseStatus{Status: "Unauthorized", Code: 401})
	}
}

// enrichement function for messages To contain name and lastname!
// just  - it does not work yet
func Enrich(uuids []gocql.UUID) map[string]string {
	if len(uuids) > 0 {
		names := map[string]string{}
		m := map[string]interface{}{}
		query := "SELECT id, name, lastname FROM users WHERE id IN ?"
		iterator := Cass.Session.Query(query, uuids).Iter()
		for iterator.MapScan(m) {
			fmt.Println("m", m)
			userID := m["id"].(gocql.UUID)
			names[userID.String()] = fmt.Sprintf("%s %s", m["name"].(string), m["lastname"].(string))
			m = map[string]interface{}{}
		}
		fmt.Printf(Cass.LogTime()+", Returning.. %s", names)
		return names
	}
	return map[string]string{}
}
