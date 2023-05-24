package utils

import (
	"github.com/streadway/amqp"
	"strings"
)

const url = "amqp://guest:guest@192.168.18.3:5672/"

func GetChannel() *amqp.Channel {
	connection, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	return channel
}

func BodyFrom(args []string) string {
	var s string
	if len(args) < 2 || args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args, " ")
	}
	return s
}
