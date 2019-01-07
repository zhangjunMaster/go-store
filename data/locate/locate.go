package locate

import (
	"fmt"
	"go-store/rabbitmq"
	"os"
	"strconv"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return os.IsNotExist(err)
}

// 绑定dataServers exchange
// 接收消息
// 判断该文件是否在本地
// 将消息发送到abbitmq
func StartLocate() {
	fmt.Println("RABBITMQ_SERVER:", os.Getenv("RABBITMQ_SERVER"))
	q, err := rabbitmq.NewRabbitMQ(os.Getenv("RABBITMQ_SERVER"))
	if err != nil {
		panic(err)
	}
	q.Bind("dataServers")
	c, _ := q.Consume()
	for msg := range c {
		object, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}
		if Locate(os.Getenv("STORAGE_ROOT") + "/objecs/" + object) {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS_DATA"))
		}
	}
}
