package main

import (
	"fmt"
	"go-store/api/heartbeat"
	"go-store/api/locate"
	"go-store/api/objects"
	"log"
	"net/http"
	"os"
)

func init() {
	os.Setenv("STORAGE_ROOT", "/Users/jun/store_dir")
	// api的启动监控端口，9200
	os.Setenv("LISTEN_ADDRESS_API", "127.0.0.1:9101")
	//rabbitmq的服务端口5672
	os.Setenv("RABBITMQ_SERVER", "amqp://guest:guest@localhost:5672")
}

func main() {
	go heartbeat.ListenHeartBeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	fmt.Println("api:", os.Getenv("LISTEN_ADDRESS_API"))
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS_API"), nil))
}
