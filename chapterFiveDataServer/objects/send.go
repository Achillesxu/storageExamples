// Package objects
// Time    : 2022/7/13 17:26
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package objects

import (
	"io"
	"os"
)

func sendFile(w io.Writer, file string) {
	f, _ := os.Open(file)
	defer f.Close()
	io.Copy(w, f)
}
