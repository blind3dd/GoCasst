package Auth

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func AuthCheck(w http.ResponseWriter, r *http.Request) bool {
	str := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(str) != 2 {
		return false
	}

	key, err := base64.StdEncoding.DecodeString(str[1])
	if err != nil {
		return false
	}

	value := strings.SplitN(string(key), ":", 2)
	if len(value) != 2 {
		return false
	}

	return value[0] == "user" && value[1] == "pass"
}
