package rest

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

func RunDynamicServer() {
	log.Println("Dynamic Server was started!")

	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/home/user/GolangProjects/src/github.com/yevchuk-kostiantyn/TestKnowledge/view"))))
	log.Fatal(http.ListenAndServe(":1997", router))
}