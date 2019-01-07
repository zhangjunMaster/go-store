package main

// 在本地取文件

import (
	"fmt"
	"go-store/data/heartbeat"
	"go-store/data/locate"
	"go-store/data/objects"
	"log"
	"net/http"
	"os"
)

func init() {
	os.Setenv("STORAGE_ROOT", "/Users/jun/store_dir")
	// data的启动监控端口，9100
	os.Setenv("LISTEN_ADDRESS_DATA", "127.0.0.1:9100")
	//rabbitmq的服务端口5672
	os.Setenv("RABBITMQ_SERVER", "amqp://guest:guest@localhost:5672")
}

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	fmt.Println("data:", os.Getenv("LISTEN_ADDRESS_DATA"))
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS_DATA"), nil))
}
