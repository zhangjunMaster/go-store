package heartbeat

import (
	"fmt"
	"go-store/rabbitmq"
	"os"
	"time"
)

func StartHeartbeat() {
	q, err := rabbitmq.NewRabbitMQ(os.Getenv("RABBITMQ_SERVER"))
	fmt.Println("RABBITMQ_SERVER:", os.Getenv("RABBITMQ_SERVER"))
	if err != nil {
		panic(err)
	}
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS_DATA"))
		time.Sleep(5 * time.Second)
	}
}
