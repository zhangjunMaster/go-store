package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	fmt.Println("m:", m)
	if m == http.MethodPut {
		fmt.Println("这里put")
		put(w, r)
		return
	}
	if m == http.MethodGet {
		fmt.Println("这里")
		get(w, r)
		return
	}
}

// func Split(s, sep string) []string
// func (u *URL) EscapedPath() string
func put(w http.ResponseWriter, r *http.Request) {
	dirPath := os.Getenv("STORAGE_ROOT") + "objects/"
	os.MkdirAll(dirPath, os.ModePerm)
	f, err := os.Create(dirPath + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	// func Copy(dst Writer, src Reader) (written int64, err error)
	io.Copy(f, r.Body)
}

func get(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(os.Getenv("STORAGE_ROOT") + "objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
