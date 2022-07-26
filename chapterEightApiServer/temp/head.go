// Package temp
// Time    : 2022/7/25 21:57
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package temp

import (
	"achilles/common/rs"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func head(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.URL.EscapedPath(), "/")[2]
	stream, e := rs.NewRSResumablePutStreamFromToken(token)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	current := stream.CurrentSize()
	if current == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("content-length", fmt.Sprintf("%d", current))
}
