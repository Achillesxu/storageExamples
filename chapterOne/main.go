// Package GoDistributeStorage
// Time    : 2022/6/22 09:17
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func HandleF(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func put(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create("/Users/achilles_xushy/Desktop" + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		log.Println(errors.Wrap(err, "create file failed"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(errors.Wrap(err, "close file failed"))
		}
	}(f)
	_, _ = io.Copy(f, r.Body)
}

func get(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("/Users/achilles_xushy/Desktop" + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		log.Println(errors.Wrap(err, "open file failed"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(errors.Wrap(err, "close file failed"))
		}
	}(f)
	_, _ = io.Copy(w, f)
}

func main() {
	http.HandleFunc("/objects/", HandleF)
	log.Fatal(http.ListenAndServe(":8601", nil))
}
