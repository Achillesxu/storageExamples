// Package objects
// Time    : 2022/6/27 23:14
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package objects

import (
	"log"
	"net/http"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	c, e := storeObject(r.Body, object)
	if e != nil {
		log.Println(e)
	}
	w.WriteHeader(c)
}
