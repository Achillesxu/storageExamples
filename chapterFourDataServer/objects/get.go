// Package objects
// Time    : 2022/7/12 08:08
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package objects

import (
	"net/http"
	"strings"
)

func get(w http.ResponseWriter, r *http.Request) {
	file := getFile(strings.Split(r.URL.EscapedPath(), "/")[2])
	if file == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	sendFile(w, file)
}
