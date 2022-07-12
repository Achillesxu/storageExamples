// Package chapterFourDataServer
// Time    : 2022/7/12 08:02
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"achilles/chapterFourDataServer/heartbeat"
	"achilles/chapterFourDataServer/locate"
	"achilles/chapterFourDataServer/objects"
	"achilles/chapterFourDataServer/temp"
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
