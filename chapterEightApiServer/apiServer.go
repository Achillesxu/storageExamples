// Package chapterEightApiServer
// Time    : 2022/7/25 21:45
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"achilles/chapterEightApiServer/heartbeat"
	"achilles/chapterEightApiServer/locate"
	"achilles/chapterEightApiServer/objects"
	"achilles/chapterEightApiServer/temp"
	"achilles/chapterEightApiServer/versions"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
