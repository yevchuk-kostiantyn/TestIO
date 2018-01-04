package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/yevchuk-kostiantyn/TestIO/database"
	"github.com/yevchuk-kostiantyn/TestIO/rest"
)

func main() {
	log.Infoln("Welcome to TestIO App log!")

	database.SaveAdmin()
	err := rest.RunDynamicServer()

	if err != nil {
		log.Errorln("RunDynamicServer()", err)
		return
	}
}
