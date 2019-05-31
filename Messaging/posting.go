package messaging

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blind3dd/gocasst/cass"

	"github.com/blind3dd/gocasst/auth"

	"github.com/gocql/gocql"
)

func PostTheMessage(w http.ResponseWriter, r *http.Request) {
	var posted bool = false
	var UUID gocql.UUID

	var errs []string
	message, errs := ProcessMessageForm(r)

	if len(errs) == 0 && auth.AuthCheck(w, r) {
		fmt.Println(cass.LogTime() + ", Creating a message..")
		UUID = gocql.TimeUUID()
		if err := cass.Session.Query(`
		INSERT INTO messages (id, user_id, user_full_name, message) VALUES (?, ?, ?, ?)`,
			UUID, message.UserID, message.UserFullName, message.Message).Exec(); err != nil {
			errs = append(errs, err.Error())
		} else {
			posted = true
		}
	} else {
		json.NewEncoder(w).Encode(ResponseStatus{Status: "Unauthorized", Code: 401})
		fmt.Println(cass.LogTime() + ", Post Message Request - Unauthorized")
	}

	if posted {
		fmt.Println(cass.LogTime()+", id", UUID)
		json.NewEncoder(w).Encode(NewMessageResponse{ID: UUID})
	} else {
		fmt.Println(cass.LogTime()+", Response Errors: ", errs)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
