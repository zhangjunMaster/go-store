package rabbitmq

import "github.com/streadway/amqp"

// 在rabbitmq中，exchange有4个类型：direct，topic，fanout，header。
// exchange的类型对应不同的队列和路由规则。
type RabbitMQ struct {
	channel  *amqp.Channel
	Name     string
	exchange string
}

// s server地址
func NewRabbitMQ(s string) (*RabbitMQ, error) {
	// rabbitmq建立连接
	conn, err := amqp.Dial(s)
	if err != nil {
		return nil, err
	}
	// 所有的操作都在channel中
	// Most operations happen on a channel.  If any error is returned on a
	// channel, the channel will no longer be valid, throw it away and try with
	// a different channel.  If you use many channels, it's useful for the
	// server to
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	// 生成queue
	// QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args Table) (Queue, error)
	q, err := channel.QueueDeclare(
		"",    // name
		false, // durable,持久化
		true,  // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}
	// new只是生成一个该类型的指针
	// new之后需要一个指针来指向这个结构
	// make会分配结构空间及其附属空间，并完成其间的指针初始化
	mq := new(RabbitMQ)
	mq.channel = channel
	mq.Name = q.Name
	// 没有 exchane
	return mq, nil
}
