package DB

import "log"

func SaveNewUser(firstName string, lastName string, email string, password string, position string) bool {
	client, err := RunDBConnection()

	if err != nil {
		log.Println("DB Error | RunDBConnection():", err)
		return false
	}

	key := email

	OK, err := client.HMSet(key, "password", password, "first_name", firstName, "last_name", lastName,
		"position", position)

	if OK != "OK" {
		log.Println("HMSet response is not OK")
		return false
	}

	if err != nil {
		log.Println("DB Error | HMSet(): ", err)
		return false
	}

	result, err := client.SAdd(position, key)
	if result == 1{
		return true
	} else if result == 0 {
		log.Println("The email already exists")
		return false
	} else {
		return false
	}
	return true
}
