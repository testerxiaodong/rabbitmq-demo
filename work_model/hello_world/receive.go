package main

import (
	"fmt"
	"rabbitmq-demo/utils"
)

func main() {
	channel := utils.GetChannel()
	// 使用默认交换机
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
	// 消费消息
	deliveries, err := channel.Consume(queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		panic(err)
	}
	for delivery := range deliveries {
		fmt.Println(string(delivery.Body))
		delivery.Ack(false)
	}
}
