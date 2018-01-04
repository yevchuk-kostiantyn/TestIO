package rest

import (
	log "github.com/sirupsen/logrus"
	"regexp"
)

func ValidateInfo(info string) bool {
	if len(info) < 2 {
		return false
	} else {
		return true
	}
}

func ValidateEmail(email string) bool {
	regularExp := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	result, err := regexp.MatchString(regularExp, email)
	if err != nil {
		log.Errorln("ValidateEmail() MatchString")
	}

	return result
}

func ValidatePassword(password string) bool {
	if len(password) < 6 {
		return false
	} else {
		return true
	}
}
