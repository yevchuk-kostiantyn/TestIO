package DB

import (
	"menteslibres.net/gosexy/redis"
	"log"
)

func RunDBConnection() (*redis.Client, error) {
	var client *redis.Client
	client = redis.New()

	err := client.Connect("127.0.0.1", 6379)
	if err != nil {
		log.Println("DB Error 1 | Connect(): ", err)
	}
	return client, err
}