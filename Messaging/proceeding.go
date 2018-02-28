package Messaging

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gocql/gocql"
)

// func ShowTime(Time uint64) string {
// 	Time = uint64(time.Now().Unix()) * 1000
// 	return strconv.FormatUint(Time, 10)
// }

func ProcessMessageForm(r *http.Request) (Message, []string) {
	var message Message
	var err error
	var errs []string
	var errStr, UserIDString string

	message.UserFullName, errStr = processFormItem(r, "user_full_name")
	errs = appendError(errs, errStr)
	message.Message, errStr = processFormItem(r, "message")
	errs = appendError(errs, errStr)

	// static static static
	UserIDString, errStr = processFormItem(r, "user_id")
	errs = appendError(errs, errStr)
	if len(errStr) != 0 {
		errs = append(errs, errStr)
	} else {
		message.UserID, err = gocql.ParseUUID(UserIDString)
		if err != nil {
			fmt.Println((strconv.FormatUint((uint64(time.Now().Unix())*1000), 10))+" Improper UUID: ", UserIDString)
			errs = append(errs, "Improper User's UUID")
		}
	}
	return message, errs
}

func processFormItem(r *http.Request, item string) (string, string) {
	itemD := r.PostFormValue(item)
	if len(itemD) == 0 {
		return "", "Missing '" + item + "'"
	}
	return itemD, ""
}

func appendError(errs []string, errStr string) []string {
	if len(errStr) > 0 {
		errs = append(errs, errStr)
	}
	return errs
}
