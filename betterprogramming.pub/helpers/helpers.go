package helpers

import (
	"log"
	"strings"
)

func CheckUserPass(username, password string) bool {
	userpass := make(map[string]string)
	userpass["hello"] = "itsme"
	userpass["john"] = "doe"

	log.Println("checkUserPass", username, password, userpass)

	if val, ok := userpass[username]; ok {
		log.Println(val, ok)
		if val == password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}