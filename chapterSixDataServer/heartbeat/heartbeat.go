// Package heartbeat
// Time    : 2022/7/13 20:40
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package heartbeat

import (
	"achilles/common/rabbitmq"
	"os"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}
