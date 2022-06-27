// Package objects
// Time    : 2022/6/27 23:14
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package objects

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func get(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	stream, e := getStream(object)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, _ = io.Copy(w, stream)

}
