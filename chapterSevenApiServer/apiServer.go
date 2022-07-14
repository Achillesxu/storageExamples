// Package chapterSevenApiServer
// Time    : 2022/7/14 09:58
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"achilles/chapterSevenApiServer/heartbeat"
	"achilles/chapterSevenApiServer/locate"
	"achilles/chapterSevenApiServer/objects"
	"achilles/chapterSevenApiServer/temp"
	"achilles/chapterSevenApiServer/versions"
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
