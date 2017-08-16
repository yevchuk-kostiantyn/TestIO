package rest

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"net/smtp"
	"github.com/yevchuk-kostiantyn/TestKnowledge"
	"encoding/json"
	"github.com/yevchuk-kostiantyn/TestKnowledge/DB"
	"strings"
)

func RunDynamicServer() {
	log.Println("Dynamic Server was started!")

	router := mux.NewRouter()

	router.HandleFunc("/login", checkCredentials).Methods("PATCH")
	router.HandleFunc("/signup", getNewUser).Methods("PATCH")

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/home/user/GolangProjects/src/github.com/yevchuk-kostiantyn/TestKnowledge/view"))))
	log.Fatal(http.ListenAndServe(":1997", router))
}

func checkCredentials(w http.ResponseWriter, r *http.Request) {
	var credentials TestKnowledge.LoginCredentials

	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		log.Println("checkCredentials(): Decode", err)
	}

	entered_username := credentials.Username
	entered_password := credentials.Password

	log.Println("Username:", entered_username)
	log.Println("Password:", entered_password)

	client, err := DB.RunDBConnection()

	if err != nil {
		log.Println("DB Error 1 | RunDBConnection():", err)
	}

	user_exists, err := client.Exists(entered_username)

	if err != nil {
		log.Println("DB Error 2 | Exists():", err)
	}

	correct_password, err := client.HMGet(entered_username, "password")

	if err != nil {
		log.Println("DB Error 3 | HMGet():", err)
	}

	if user_exists || entered_password == strings.Join(correct_password, "") {
		var response TestKnowledge.Response
		response.Position, err = client.HMGet(entered_username, "position")
		if err != nil {
			log.Println("DB Error 4 | HMGet():", err)
		}
		response_position_string := strings.Join(response.Position, "")
		log.Println("Position:", response_position_string)
		json.NewEncoder(w).Encode(response_position_string)
	}
}

func getNewUser(_ http.ResponseWriter, r *http.Request) {
	var newUserInfo TestKnowledge.NewUser

	err := json.NewDecoder(r.Body).Decode(&newUserInfo)

	if err != nil {
		log.Println("checkCredentials(): Decode", err)
	}

	firstName := newUserInfo.FirstName
	lastName := newUserInfo.LastName
	email := newUserInfo.Email
	password := newUserInfo.Password
	position := newUserInfo.Position

	log.Println("New User:")
	log.Println("Name:", firstName, lastName)
	log.Println("Email:", email)
	log.Println("Password:", password)
	log.Println("Position:", position)

	saveToDB(firstName, lastName, email, password, position)
}

func saveToDB(firstName string, lastName string, email string, password string, position string) {
	client, err := DB.RunDBConnection()

	if err != nil {
		log.Println("DB Error 5 | RunDBConnection():", err)
	}

	key := email

	OK, err := client.HMSet(key, "password", password, "first_name", firstName, "last_name", lastName,
	"position", position)

	if OK != "OK" {
		log.Println("HMSet response is not OK")
	}

	if err != nil {
		log.Println("DB Error 6 | HMSet(): ", err)
	}

	sendEmail(firstName, lastName, email, password, position)
}

func sendEmail(firstName string, lastName string, email string, password string, position string) {
	from := "kostiantyn.yevchuk@nure.ua"
	admin_password := "yewchuk97"

	body := "Dear " + firstName +" " + lastName +"," + "\n" + "You were successfully signed up on TestIO." + "\n" +
		"To login enter the email and password" +
		" you used." + "\n" + "Thank you!" + "\n" + "Have a great day!" + "\n" + "TestIO Team"

	msg := "From: " + from + "\r\n" +
			"To: " + email + "\r\n" +
			"Subject: Your messages subject" + "\r\n\r\n" + body + "\r\n"
	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, admin_password, "smtp.gmail.com"), from, []string{email}, []byte(msg))
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	log.Print("Email to", email, "was successfully sent")
}