package main

import (
	"go-store/data/heartbeat"
	"go-store/data/locate"
	"go-store/data/objects"
	"log"
	"net/http"
	"os"
)

func init() {
	os.Setenv("STORAGE_ROOT", "/tmp/")
	// data的启动监控端口，9100
	os.Setenv("LISTEN_ADDRES", "127.0.0.1:9100")
	//rabbitmq的服务端口5672
	os.Setenv("RABBITMQ_SERVER", "amqp://guest:guest@192.168.3.167:5672")
}

func main() {
	go heartbeat.StartHeartBeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRES"), nil))
}
