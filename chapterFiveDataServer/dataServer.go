// Package chapterFiveDataServer
// Time    : 2022/7/13 17:22
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package chapterFiveDataServer

import (
	"achilles/chapterFiveDataServer/heartbeat"
	"achilles/chapterFiveDataServer/objects"
	"achilles/chapterFourDataServer/locate"
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
