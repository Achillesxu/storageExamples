// Package chapterSixApiServer
// Time    : 2022/7/13 20:04
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"achilles/chapterSixApiServer/heartbeat"
	"achilles/chapterSixApiServer/locate"
	"achilles/chapterSixApiServer/objects"
	"achilles/chapterSixApiServer/temp"
	"achilles/chapterSixApiServer/versions"
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
