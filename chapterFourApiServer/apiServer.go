// Package chapterFourApiServer
// Time    : 2022/7/12 07:24
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"achilles/chapterFourApiServer/heartbeat"
	"achilles/chapterFourApiServer/locate"
	"achilles/chapterFourApiServer/objects"
	"achilles/chapterFourApiServer/versions"
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
