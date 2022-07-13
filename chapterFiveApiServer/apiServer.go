// Package chapterFiveApiServer
// Time    : 2022/7/12 20:51
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"achilles/chapterFiveApiServer/heartbeat"
	"achilles/chapterFiveApiServer/locate"
	"achilles/chapterFiveApiServer/objects"
	"achilles/chapterFiveApiServer/versions"
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
