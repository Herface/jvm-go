package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
)

// ZipEntry .zip .jar类库
type ZipEntry struct {
	absPath string
	reader  *zip.Reader
}

func (this *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	reader := this.reader
	for _, file := range reader.File {
		if file.Name == className {
			open, err := file.Open()
			if err != nil {
				panic(err)
			}
			all, err := ioutil.ReadAll(open)
			return all, this, err
		}
	}
	return nil, nil, errors.New("could not find class")
}

func (this *ZipEntry) String() string {
	return this.absPath
}
