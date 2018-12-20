package heartbeat

import (
	"go-store/data/rabbitmq"
	"log"
	"os"
	"time"
)

//  q.Publish 中只有exchange,因此会将信息发送到所有的queue
func StartHeartBeat() {
	q, err := rabbitmq.NewRabbitMQ(os.Getenv("RABBITMQ_SERVER"))
	if err != nil {
		panic(err)
	}
	defer q.Close()
	for {
		// 每5秒向apiServers这个exchange发送LISTEN_ADDRES地址
		err := q.Publish("apiServers", os.Getenv("LISTEN_ADDRES"))
		if err != nil {
			log.Println(err)
		}
		time.Sleep(5 * time.Second)
	}
}
