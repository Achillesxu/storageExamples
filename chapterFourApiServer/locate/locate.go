// Package locate
// Time    : 2022/7/12 07:29
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package locate

import (
	"achilles/common/rabbitmq"
	"os"
	"strconv"
	"time"
)

func Locate(name string) string {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

func Exist(name string) bool {
	return Locate(name) != ""
}
