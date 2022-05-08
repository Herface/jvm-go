package classpath

import (
	"io/ioutil"
	"path"
)

type DirEntry struct {
	absPath string
}

func (this *DirEntry) ReadClass(className string) ([]byte, Entry, error) {
	filePath := path.Join(this.absPath, className)
	data, err := ioutil.ReadFile(filePath)
	return data, this, err
}
func (this *DirEntry) String() string {
	return this.absPath
}
