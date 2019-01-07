package locate

import (
	"encoding/json"
	"go-store/rabbitmq"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// 接收定位消息，定位对象，发送反馈消息
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	info, err := Locate(strings.Split(r.URL.EscapedPath(), "/")[2])
	if len(info) == 0 || err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(info)
	w.Write(b)
}

// func Stat(name string) (FileInfo, error)
// func IsNotExist(err error) bool
func Locate(name string) (string, error) {
	q, err := rabbitmq.NewRabbitMQ(os.Getenv("RABBITMQ_SERVER"))
	if err != nil {
		return "", err
	}
	// 向 exchange publish
	q.Publish("dataServes", name)
	c, err := q.Consume()
	if err != nil {
		return "", err
	}
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s, nil

}

func Exist(name string) bool {
	s, err := Locate(name)
	if err != nil || s == "" {
		return false
	}
	return true
}
