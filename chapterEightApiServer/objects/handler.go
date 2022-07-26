// Package objects
// Time    : 2022/7/25 21:53
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package objects

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPost {
		post(w, r)
		return
	}
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	if m == http.MethodDelete {
		del(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
