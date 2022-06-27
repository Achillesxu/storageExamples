// Package chapterThreeDataServer
// Time    : 2022/6/28 03:40
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"achilles/chapterThreeDataServer/heartbeat"
	"achilles/chapterThreeDataServer/locate"
	"achilles/chapterThreeDataServer/objects"
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
