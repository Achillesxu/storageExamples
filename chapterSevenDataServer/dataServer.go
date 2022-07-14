// Package chapterSevenDataServer
// Time    : 2022/7/14 10:02
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"achilles/chapterFiveDataServer/temp"
	"achilles/chapterSevenDataServer/heartbeat"
	"achilles/chapterSevenDataServer/locate"
	"achilles/chapterSevenDataServer/objects"
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
