package rest

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"net/smtp"
	"github.com/yevchuk-kostiantyn/TestKnowledge"
	"encoding/json"
	"github.com/yevchuk-kostiantyn/TestKnowledge/DB"
)

func RunDynamicServer() {
	log.Println("Dynamic Server was started!")

	router := mux.NewRouter()

	router.HandleFunc("/login", checkEnteredCredentials).Methods("PATCH")
	router.HandleFunc("/signup", getNewUser).Methods("PATCH")

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/home/user/GolangProjects/src/github.com/yevchuk-kostiantyn/TestKnowledge/view"))))
	log.Fatal(http.ListenAndServe(":1997", router))
}

func checkEnteredCredentials(w http.ResponseWriter, r *http.Request) {
	var credentials TestKnowledge.LoginCredentials

	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		log.Println("checkCredentials(): Decode", err)
	}

	entered_username := credentials.Username
	entered_password := credentials.Password

	log.Println("Username:", entered_username)
	log.Println("Password:", entered_password)

	user_exists := DB.UserExists(entered_username)
	is_password_correct := DB.IsPasswordCorrect(entered_username, entered_password)

	if user_exists && is_password_correct {
		var response TestKnowledge.Response
		response.Position = DB.GetUserPosition(entered_username)
		log.Println("Position:", response.Position)
		json.NewEncoder(w).Encode(response.Position)
	} else if !user_exists {
		log.Println("User not exists")
	} else if !is_password_correct {
		log.Println("Incorrect Password")
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

	successful_save := DB.SaveNewUser(firstName, lastName, email, password, position)

	if successful_save {
		sendEmail(firstName, lastName, email, position)
	} else {
		log.Println("Save to DB was not successful")
	}
}

func sendEmail(firstName string, lastName string, email string, position string) {
	from := "kostiantyn.yevchuk@nure.ua"
	admin_password := "yewchuk97"

	body := "Dear " + firstName +" " + lastName +"," + "\n" + "You were successfully signed up on" +
		" TestIO as " + position + "." + "\n" + "To login enter the email and password" +
		" you used." + "\n" + "Thank you!" + "\n" + "Have a great day!" + "\n" + "TestIO Team"

	msg := "From: " + from + "\r\n" +
			"To: " + email + "\r\n" +
			"Subject: Your messages subject" + "\r\n\r\n" + body + "\r\n"
	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, admin_password, "smtp.gmail.com"), from, []string{email}, []byte(msg))
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	log.Print("Email to ", email, " was successfully sent")
}