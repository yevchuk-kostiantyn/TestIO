package database

import (
	. "github.com/smartystreets/goconvey/convey"
	"menteslibres.net/gosexy/redis"
	"testing"
)

func TestRunDBConnection(t *testing.T) {
	var client *redis.Client
	client = redis.New()

	dbHost := "127.0.0.1"
	var dbPort uint
	dbPort = 6379
	Convey("Check redis client connection"+dbHost+":"+string(dbPort)+". Should be without error ", t, func() {
		err := client.Connect(dbHost, dbPort)
		defer client.Close()
		So(err, ShouldBeNil)
	})
}

func TestFlushall(t *testing.T) {
	expectedOutput := Flushall()

	if expectedOutput != nil {
		t.Error("Expected nil, got ", expectedOutput)
	}
}
