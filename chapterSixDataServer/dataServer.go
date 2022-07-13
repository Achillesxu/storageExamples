// Package chapterSixDataServer
// Time    : 2022/7/13 20:39
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"achilles/chapterSixDataServer/heartbeat"
	"achilles/chapterSixDataServer/locate"
	"achilles/chapterSixDataServer/objects"
	"achilles/chapterSixDataServer/temp"
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
