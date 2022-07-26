// Package temp
// Time    : 2022/7/25 22:06
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package temp

import (
	"net/http"
	"os"
	"strings"
)

func del(w http.ResponseWriter, r *http.Request) {
	uuid := strings.Split(r.URL.EscapedPath(), "/")[2]
	infoFile := os.Getenv("STORAGE_ROOT") + "/temp/" + uuid
	datFile := infoFile + ".dat"
	os.Remove(infoFile)
	os.Remove(datFile)
}
