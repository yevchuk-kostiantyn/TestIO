package database

import log "github.com/sirupsen/logrus"

func SaveNewUser(firstName string, lastName string, email string, password string, position string) bool {
	client, err := RunDBConnection()
	if err != nil {
		log.Errorln("database | RunDBConnection():", err)
		return false
	}

	key := email

	OK, err := client.HMSet(key, "password", password, "first_name", firstName, "last_name", lastName,
		"position", position)

	if err != nil {
		log.Errorln("database | HMSet(): ", err)
		return false
	}

	if OK != "OK" {
		log.Warningln("HMSet response is not OK")
		return false
	}

	result, err := client.SAdd(position, key)
	if err != nil {
		log.Errorln("database | SAdd()")
		return false
	}

	if result == 1{
		return true
	} else if result == 0 {
		log.Warningln("The email already exists")
		return false
	} else {
		return false
	}
	return true

}

func SaveAdmin() bool {
	client, err := RunDBConnection()
	if err != nil {
		log.Errorln("database | RunDBConnection():", err)
		return false
	}

	OK, err := client.HMSet("kostiantyn.yevchuk@nure.ua", "password", "yewchuk97", "first_name", "Kostiantyn", "last_name", "Yevchuk",
		"position", "admin")

	if err != nil {
		log.Errorln("database | HMSet(): ", err)
		return false
	}

	if OK != "OK" {
		log.Warningln("HMSet response is not OK")
		return false
	}
	return true
}
