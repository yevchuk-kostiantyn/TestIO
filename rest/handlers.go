package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/yevchuk-kostiantyn/TestIO/database"
	"github.com/yevchuk-kostiantyn/TestIO/models"
	"net/http"
	"net/smtp"
	"os"
)

func RunDynamicServer() error {
	log.Info("Dynamic Server was started!")

	router := mux.NewRouter()

	router.HandleFunc("/login", checkEnteredCredentials).Methods("PATCH")
	router.HandleFunc("/signup", getNewUser).Methods("PATCH")

	GOPATH := os.Getenv("GOPATH")
	viewPath := "/src/github.com/yevchuk-kostiantyn/TestIO/view/"

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(GOPATH+viewPath))))

	err := http.ListenAndServe(":1997", router)
	if err != nil {
		log.Errorln("ListenAndServe()", err)
		return err
	}
	return nil
}

func checkEnteredCredentials(w http.ResponseWriter, r *http.Request) {
	var credentials models.LoginCredentials

	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		log.Errorln("checkCredentials(): Decode", err)
	}

	entered_username := credentials.Username
	entered_password := credentials.Password

	log.Println("Username:", entered_username)
	log.Println("Password:", entered_password)

	user_exists := database.UserExists(entered_username)
	is_password_correct := database.IsPasswordCorrect(entered_username, entered_password)

	if user_exists && is_password_correct {
		var response models.Response
		response.Position = database.GetUserPosition(entered_username)
		log.Println("Position:", response.Position)
		err := json.NewEncoder(w).Encode(response.Position)
		if err != nil {
			log.Errorln("JSON Encode()")
		}
	} else if !user_exists {
		log.Warningln("User does not exist")
	} else if !is_password_correct {
		log.Warningln("Incorrect Password")
	}
}

func getNewUser(_ http.ResponseWriter, r *http.Request) {
	var newUserInfo models.NewUser

	err := json.NewDecoder(r.Body).Decode(&newUserInfo)

	if err != nil {
		log.Errorln("checkCredentials(): Decode", err)
	}

	firstName := newUserInfo.FirstName
	lastName := newUserInfo.LastName
	email := newUserInfo.Email
	password := newUserInfo.Password
	position := newUserInfo.Position

	switch {
	case !ValidateInfo(firstName):
		log.Warningln("Entered first name is invalid")
	case !ValidateInfo(lastName):
		log.Warningln("Entered last name is invalid")
	case !ValidateEmail(email):
		log.Warningln("Entered email is invalid")
	case !ValidatePassword(password):
		log.Warningln("Entered password is invalid")
	case ValidateInfo(firstName) || ValidateInfo(lastName) || ValidateEmail(email) ||
		ValidatePassword(password):
		log.Println("New User:")
		log.Println("Name:", firstName, lastName)
		log.Println("Email:", email)
		log.Println("Password:", password)
		log.Println("Position:", position)

		successful_save := database.SaveNewUser(firstName, lastName, email, password, position)

		if successful_save {
			sendEmail(firstName, lastName, email, position)
		} else {
			log.Println("Save to database was not successful")
		}
	}
}

func sendEmail(firstName string, lastName string, email string, position string) {
	from := "kostiantyn.yevchuk@nure.ua"
	admin_password := "yewchuk97"

	body := "Dear " + firstName + " " + lastName + "," + "\n" + "You were successfully signed up on" +
		" TestIO as " + position + "." + "\n" + "To login enter the email and password" +
		" you used." + "\n" + "Thank you!" + "\n" + "Have a great day!" + "\n" + "TestIO Team"

	msg := "From: " + from + "\r\n" +
		"To: " + email + "\r\n" +
		"Subject: Your messages subject" + "\r\n\r\n" + body + "\r\n"
	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, admin_password, "smtp.gmail.com"), from, []string{email}, []byte(msg))
	if err != nil {
		log.Errorf("Error: %s", err)
		return
	}

	log.Print("Email to ", email, " was successfully sent")
}
