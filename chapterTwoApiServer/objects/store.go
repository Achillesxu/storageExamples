// Package objects
// Time    : 2022/6/27 23:14
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package objects

import (
	"io"
	"net/http"
)

func storeObject(r io.Reader, object string) (int, error) {
	stream, e := putStream(object)
	if e != nil {
		return http.StatusServiceUnavailable, e
	}

	_, _ = io.Copy(stream, r)
	e = stream.Close()
	if e != nil {
		return http.StatusInternalServerError, e
	}
	return http.StatusOK, nil
}
