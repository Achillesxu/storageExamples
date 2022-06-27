// Package chapterTwoDataServer
// Time    : 2022/6/27 23:00
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"achilles/chapterTwoDataServer/heartbeat"
	"achilles/chapterTwoDataServer/locate"
	"achilles/chapterTwoDataServer/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
