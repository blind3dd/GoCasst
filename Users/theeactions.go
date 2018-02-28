package Users

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func processFormItem(r *http.Request, item string) (string, string) {
	itemData := r.PostFormValue(item)
	if len(itemData) == 0 {
		return "", "Missing '" + item + "'"
	}
	return itemData, ""
}

func ProcessUserFrom(r *http.Request) (User, []string) {
	var user User
	var err error
	var errs []string
	var errStr, birthYearStr string

	user.Name, errStr = processFormItem(r, "name")
	errs = appendError(errs, errStr)
	user.LastName, errStr = processFormItem(r, "lastname")
	errs = appendError(errs, errStr)
	user.Email, errStr = processFormItem(r, "email")
	errs = appendError(errs, errStr)
	user.City, errStr = processFormItem(r, "city")
	errs = appendError(errs, errStr)

	birthYearStr, errStr = processFormItem(r, "birthyear")
	if len(errStr) != 0 {
		errs = append(errs, errStr)
	} else {
		user.BirthYear, err = strconv.Atoi(birthYearStr)
		if err != nil {
			fmt.Println((strconv.FormatUint((uint64(time.Now().Unix())*1000), 10))+" Improper birthyear value: ", birthYearStr)
			errs = append(errs, "'birthyear' has to integer value, not a string!")
		}
	}
	return user, errs
}

func appendError(errs []string, errStr string) []string {
	if len(errStr) > 0 {
		errs = append(errs, errStr)
	}
	return errs
}
