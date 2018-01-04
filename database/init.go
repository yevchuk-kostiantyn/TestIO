package database

import (
	log "github.com/sirupsen/logrus"
	"menteslibres.net/gosexy/redis"
)

func RunDBConnection() (*redis.Client, error) {
	var client *redis.Client
	client = redis.New()

	databaseHost := "127.0.0.1"
	databasePort := 6379

	err := client.Connect(databaseHost, uint(databasePort))
	if err != nil {
		log.Errorln("database | Connect(): ", err)
	}

	return client, err
}

func Flushall() error {
	client, _ := RunDBConnection()
	_, err := client.FlushAll()

	if err != nil {
		return err
	} else {
		return nil
	}
}
