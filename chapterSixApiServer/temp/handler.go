// Package temp
// Time    : 2022/7/13 20:10
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package temp

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodHead {
		head(w, r)
		return
	}
	if m == http.MethodPut {
		put(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
