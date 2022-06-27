// Package chapterTwoApiServer
// Time    : 2022/6/27 22:59
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"achilles/chapterTwoApiServer/heartbeat"
	"achilles/chapterTwoApiServer/locate"
	"achilles/chapterTwoApiServer/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
