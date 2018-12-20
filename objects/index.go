package objects

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	status, err := storeObject(r.Body, object)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(status)
}

func storeObject(body io.ReadCloser, object string) (int, error) {

	stream, err := putStream(object)

}

func putStream(object string) *objectstream.PutSteram {

}
