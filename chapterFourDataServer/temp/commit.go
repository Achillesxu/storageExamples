// Package temp
// Time    : 2022/7/12 08:11
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package temp

import (
	"achilles/chapterFourDataServer/locate"
	"os"
)

func commitTempObject(datFile string, tempinfo *tempInfo) {
	os.Rename(datFile, os.Getenv("STORAGE_ROOT")+"/objects/"+tempinfo.Name)
	locate.Add(tempinfo.Name)
}
