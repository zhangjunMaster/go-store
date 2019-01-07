package objects

import (
	"fmt"
	"go-store/api/heartbeat"
	"go-store/api/locate"
	"go-store/api/objectstream/getstream"
	objectstream "go-store/api/objectstream/putstream"
	"io"
	"log"
	"net/http"
	"strings"
)

// 本地处理数据
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	m := r.Method
	if m == http.MethodGet {
		get(w, r)
		return
	}

	if m == http.MethodPut {
		put(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	status, err := storeObject(r.Body, object)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(status)
}

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any data server")
	}
	return objectstream.NewPutStream(server, object), nil
}

func storeObject(r io.Reader, object string) (int, error) {
	stream, err := putStream(object)
	if err != nil {
		return http.StatusServiceUnavailable, err
	}
	io.Copy(stream, r)
	err = stream.Close()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func get(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	stream, err := getStream(object)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	io.Copy(w, stream)
}

func getStream(object string) (io.Reader, error) {
	server, err := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("not found")
	}

	r, err := getstream.NewGetStream(server, object)
	if err != nil {
		return nil, err
	}
	return r, nil
}
