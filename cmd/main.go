package main

import (
	"log"
	"github.com/yevchuk-kostiantyn/TestKnowledge/rest"
)

func main() {
	log.Println("Welcome to TestIO App log!")
	rest.RunDynamicServer()
}