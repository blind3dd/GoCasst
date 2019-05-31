package messaging

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blind3dd/gocasst/auth"
	"github.com/blind3dd/gocasst/cass"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func GetAllMessagesRequest(w http.ResponseWriter, r *http.Request) {
	var getAll []Message
	m := map[string]interface{}{}

	q := "SELECT id, user_id, user_full_name, message from messages"
	i := cass.Session.Query(q).Iter()
	for i.MapScan(m) {
		getAll = append(getAll, Message{
			ID:           m["id"].(gocql.UUID),
			UserID:       m["user_id"].(gocql.UUID),
			UserFullName: m["user_full_name"].(string),
			Message:      m["message"].(string),
		})
		m = map[string]interface{}{}
	}
	if auth.AuthCheck(w, r) {
		fmt.Println(cass.LogTime() + ", Getting Messages list ..")
		json.NewEncoder(w).Encode(AllMessagesResponse{Messages: getAll})
	} else {
		json.NewEncoder(w).Encode(ResponseStatus{Status: "Unauthorized", Code: 401})
	}
}

func GetOneMsgRequest(w http.ResponseWriter, r *http.Request) {
	var exists bool = false
	var message Message
	var errs []string

	vars := mux.Vars(r)
	id := vars["message_uuid"]
	uuid, err := gocql.ParseUUID(id)
	// UserIDStr Mapping is in Proceeding .g o file  (Process Message Form)
	//userIdStr := vars["user_id"]
	//message.UserID = userIdStr

	if err != nil {
		errs = append(errs, err.Error())
	} else {
		m := map[string]interface{}{}
		q := "SELECT id, user_id, user_full_name, message FROM messages WHERE id=?"
		i := cass.Session.Query(q, uuid).Consistency(gocql.One).Iter()
		for i.MapScan(m) {
			exists = true
			message = Message{
				ID:           m["id"].(gocql.UUID),
				UserID:       m["user_id"].(gocql.UUID),
				UserFullName: m["user_full_name"].(string),
				Message:      m["message"].(string),
			}
		}
		if !exists {
			if !auth.AuthCheck(w, r) {
				json.NewEncoder(w).Encode(ResponseStatus{Status: "Unauthorized", Code: 401})
				return
			} else {
				json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
				errs = append(errs, "No message found")
			}
		}
	}
	if exists {
		if !auth.AuthCheck(w, r) {
			json.NewEncoder(w).Encode(ResponseStatus{Status: "Unauthorized", Code: 401})
			return
		} else {
			json.NewEncoder(w).Encode(GetMessageResponse{Message: message})
			fmt.Println(cass.LogTime()+", Getting message details:", message.ID)

		}
	}
}
