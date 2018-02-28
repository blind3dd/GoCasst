package Messaging

import (
	"encoding/json"
	"fmt"
	"net/http"

	"GoCasst/Auth"
	"GoCasst/Cass"

	"github.com/gocql/gocql"
)

func PostTheMessage(w http.ResponseWriter, r *http.Request) {
	var posted bool = false
	var UUID gocql.UUID

	var errs []string
	message, errs := ProcessMessageForm(r)

	if len(errs) == 0 && Auth.AuthCheck(w, r) {
		fmt.Println(Cass.LogTime() + ", Creating a message..")
		UUID = gocql.TimeUUID()
		if err := Cass.Session.Query(`
		INSERT INTO messages (id, user_id, user_full_name, message) VALUES (?, ?, ?, ?)`,
			UUID, message.UserID, message.UserFullName, message.Message).Exec(); err != nil {
			errs = append(errs, err.Error())
		} else {
			posted = true
		}
	} else {
		json.NewEncoder(w).Encode(ResponseStatus{Status: "Unauthorized", Code: 401})
		fmt.Println(Cass.LogTime() + ", Post Message Request - Unauthorized")
	}

	if posted {
		fmt.Println(Cass.LogTime()+", id", UUID)
		json.NewEncoder(w).Encode(NewMessageResponse{ID: UUID})
	} else {
		fmt.Println(Cass.LogTime()+", Response Errors: ", errs)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
