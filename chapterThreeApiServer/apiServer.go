// Package chapterThreeApiServer
// Time    : 2022/6/28 03:41
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"achilles/chapterThreeApiServer/heartbeat"
	"achilles/chapterThreeApiServer/locate"
	"achilles/chapterThreeApiServer/objects"
	"achilles/chapterThreeApiServer/versions"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
