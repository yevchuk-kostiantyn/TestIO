package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/yevchuk-kostiantyn/TestIO/rest"
	"github.com/yevchuk-kostiantyn/TestIO/DB"
)

func main() {
	log.Infoln("Welcome to TestIO App log!")

	DB.SaveAdmin()
	rest.RunDynamicServer()
}