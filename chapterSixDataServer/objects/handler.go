// Package objects
// Time    : 2022/7/13 20:41
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package objects

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodGet {
		get(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
