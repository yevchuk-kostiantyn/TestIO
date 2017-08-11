package rest

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/yevchuk-kostiantyn/TestKnowledge"
	"encoding/json"
)

func RunDynamicServer() {
	log.Println("Dynamic Server was started!")

	router := mux.NewRouter()

	router.HandleFunc("/patch", checkCredentials).Methods("PATCH")

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/home/user/GolangProjects/src/github.com/yevchuk-kostiantyn/TestKnowledge/view"))))
	log.Fatal(http.ListenAndServe(":1997", router))
}

func checkCredentials(_ http.ResponseWriter, r *http.Request) {
	var credentials TestKnowledge.LoginCredentials

	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		log.Println("checkCredentials(): Decode", err)
	}

	log.Println("Username:", credentials.Username)
	log.Println("Password:", credentials.Password)

}