package rest

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/yevchuk-kostiantyn/TestKnowledge"
	"encoding/json"
	"menteslibres.net/gosexy/redis"
)

func RunDynamicServer() {
	log.Println("Dynamic Server was started!")

	router := mux.NewRouter()

	router.HandleFunc("/get", checkCredentials).Methods("GET")
	router.HandleFunc("/signup", getNewUser).Methods("PATCH")

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/home/user/GolangProjects/src/github.com/yevchuk-kostiantyn/TestKnowledge/view"))))
	log.Fatal(http.ListenAndServe(":1997", router))
}

func checkCredentials(w http.ResponseWriter, r *http.Request) {
	var credentials TestKnowledge.LoginCredentials

	log.Println("Body:", r.Body)
	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		log.Println("checkCredentials(): Decode", err)
	}

	log.Println("Username:", credentials.Username)
	log.Println("Password:", credentials.Password)


	if credentials.Username == "kostiantyn.yevchuk@nure.ua" && credentials.Password == "yewchuk97" {
		var response TestKnowledge.Response
		response.Position = "admin"
		json.NewEncoder(w).Encode(response)
	}
}

func getNewUser(w http.ResponseWriter, r *http.Request) {
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

func runDBConnection() *redis.Client {
	var client *redis.Client
	client = redis.New()

	err := client.Connect("127.0.0.1", 6379)
	if err != nil {
		log.Println("Error Connect(): ", err)
	}
	return client
}

func saveToDB(firstName string, lastName string, email string, password string, position string) {
	client := runDBConnection()
	log.Println("RedisClient:", client)
}