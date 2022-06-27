// Package objects
// Time    : 2022/6/27 23:14
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package objects

import (
	"achilles/chapterTwoApiServer/locate"
	"achilles/common/objectstream"
	"fmt"
	"io"
)

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return objectstream.NewGetStream(server, object)
}
