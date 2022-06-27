// Package heartbeat
// Time    : 2022/6/27 23:01
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package heartbeat

import "math/rand"

func ChooseRandomDataServer() string {
	ds := GetDataServers()
	n := len(ds)
	if n == 0 {
		return ""
	}
	return ds[rand.Intn(n)]
}
