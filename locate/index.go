package locate

import (
	"go-store/data/rabbitmq"
	"os"
	"strconv"
)

// func Stat(name string) (FileInfo, error)
// func IsNotExist(err error) bool
func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// 定位某个文件在哪个服务器上
func StartLocate() {
	q, err := rabbitmq.NewRabbitMQ(os.Getenv("RABBITMQ_SERVER"))
	if err != nil {
		panic(err)
	}
	defer q.Close()
	// bind exchange
	q.Bind("dataServers")

	c, err := q.Consume()
	if err != nil {
		panic(err)
	}
	// Delivery is struct
	// strconv.Unquote 是反处理json的，去掉 “”
	for msg := range c {
		object, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}
		if Locate(os.Getenv("STORAGE_ROOT") + "/objects" + object) {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRES"))
		}
	}
}
