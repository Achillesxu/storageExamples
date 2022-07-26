// Package chapterEightDataServer
// Time    : 2022/7/25 22:04
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"achilles/chapterEightDataServer/heartbeat"
	"achilles/chapterEightDataServer/locate"
	"achilles/chapterEightDataServer/objects"
	"achilles/chapterEightDataServer/temp"
	"log"
	"net/http"
	"os"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
