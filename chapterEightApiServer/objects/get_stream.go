// Package objects
// Time    : 2022/7/25 21:52
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package objects

import (
	"achilles/chapterEightApiServer/heartbeat"
	"achilles/chapterEightApiServer/locate"
	"achilles/common/rs"
	"fmt"
)

func GetStream(hash string, size int64) (*rs.RSGetStream, error) {
	locateInfo := locate.Locate(hash)
	if len(locateInfo) < rs.DATA_SHARDS {
		return nil, fmt.Errorf("object %s locate fail, result %v", hash, locateInfo)
	}
	dataServers := make([]string, 0)
	if len(locateInfo) != rs.ALL_SHARDS {
		dataServers = heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS-len(locateInfo), locateInfo)
	}
	return rs.NewRSGetStream(locateInfo, dataServers, hash, size)
}
