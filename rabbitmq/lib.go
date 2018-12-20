package rabbitmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

// bind是将channel绑定到exchage上
// func (ch *Channel) QueueBind(name, key, exchange string, noWait bool, args Table) error
func (q *RabbitMQ) Bind(exchange string) {
	err := q.channel.QueueBind(
		q.Name,   //queue name
		"",       // routing key
		exchange, // exchane
		false,    //nowait
		nil,
	)
	if err != nil {
		panic(err)
	}
	// 在定义mq时未设置exchange
	q.exchange = exchange
}

// 向 queue发送消息
// func (ch *Channel) Publish(exchange, key string, mandatory, immediate bool, msg Publishing) error
// Publish sends a Publishing from the client to an exchange on the server
// exchange 和 queue 一起才会确定信息的路径
// json.Marshal是将map,struct解析成json字符串
func (q *RabbitMQ) Send(queue string, body interface{}) {
	str, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	err = q.channel.Publish(
		"",    // exchange 这里先设置为空
		queue, // 绑定key
		false,
		false,
		amqp.Publishing{
			ReplyTo: q.Name,
			Body:    []byte(str),
		},
	)
	if err != nil {
		panic(err)
	}
}

func (q *RabbitMQ) Publish(exchange string, body interface{}) (err error) {
	var str []byte
	if str, err = json.Marshal(body); err != nil {
		return
	}
	// 设置 exchange
	// queue是通道是空的，会将信息发送到所有的queue中
	err = q.channel.Publish(exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ReplyTo: q.Name,
			Body:    []byte(str),
		},
	)
	if err != nil {
		return
	}
	return
}

func (q *RabbitMQ) Close() {
	q.channel.Close()
}

// 消费
// func (ch *Channel) Consume(queue, consumer string, autoAck,exclusive, noLocal, noWait bool, args Table)
// (<-chan Delivery, error)
func (q *RabbitMQ) Consume() (c <-chan amqp.Delivery, err error) {
	if c, err = q.channel.Consume(q.Name, // 这个应该是queue，为什么是Name?
		"", true, false, false, false, nil); err != nil {
		return
	}
	return
}
