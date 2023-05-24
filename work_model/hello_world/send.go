package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"rabbitmq-demo/utils"
)

func main() {
	channel := utils.GetChannel()
	// 使用默认交换机，无需声明
	// 声明队列
	queue, err := channel.QueueDeclare("hello world",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		panic(err)
	}
	// 发布消息
	err = channel.Publish("",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte("hello world"),
			DeliveryMode: amqp.Persistent,
		})
	if err != nil {
		fmt.Println("发送消息失败")
	}
}
