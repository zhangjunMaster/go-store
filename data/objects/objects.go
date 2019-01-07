package objects

import (
	"io"
	"log"
	"net/http"
	"os"
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

func get(w http.ResponseWriter, r *http.Request) {
	filePath := os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2]
	f, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

func put(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}
