// Package objects
// Time    : 2022/6/27 23:14
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package objects

import (
	"achilles/chapterTwoApiServer/heartbeat"
	"achilles/common/objectstream"
	"fmt"
)

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return objectstream.NewPutStream(server, object), nil
}
