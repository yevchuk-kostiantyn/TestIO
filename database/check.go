package database

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

func UserExists(entered_username string) bool {
	client, err := RunDBConnection()
	if err != nil {
		log.Errorln("database | RunDBConnection():", err)
	}

	user_exists, err := client.Exists(entered_username)
	if err != nil {
		log.Errorln("database | Exists():", err)
	}

	return user_exists
}

func IsPasswordCorrect(entered_username string, entered_password string) bool {
	client, err := RunDBConnection()
	if err != nil {
		log.Errorln("database | RunDBConnection():", err)
	}

	correct_password, err := client.HMGet(entered_username, "password")
	if err != nil {
		log.Errorln("database | HMGet():", err)
	}

	if entered_password == strings.Join(correct_password, "") {
		return true
	} else {
		return false
	}
}

func GetUserPosition(entered_username string) string {
	client, err := RunDBConnection()
	if err != nil {
		log.Errorln("database | RunDBConnection():", err)
	}

	position, err := client.HMGet(entered_username, "position")
	if err != nil {
		log.Errorln("database | HMGet():", err)
	}

	return strings.Join(position, "")
}