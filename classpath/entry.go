package classpath

// 类路径 Entry
import (
	"archive/zip"
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const pathSeparator = string(os.PathSeparator)

type Entry interface {
	ReadClass(className string) ([]byte, Entry, error)
	String() string
}

type CompositeEntry []Entry

func (this CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, entry := range this {
		class, e, err := entry.ReadClass(className)
		if err == nil {
			return class, e, err
		}
	}
	return nil, nil, errors.New("no class found")
}
func (this CompositeEntry) String() string {
	paths := make([]string, 0)
	for _, entry := range this {
		paths = append(paths, entry.String())
	}
	return strings.Join(paths, pathSeparator)
}

func newEntry(path string) Entry {
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.Contains(path, pathSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, "zip") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}

func newZipEntry(path string) Entry {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	// 只读一次 载入内存
	data, err := ioutil.ReadFile(abs)
	if err != nil {
		panic(err)
	}
	reader := bytes.NewReader(data)
	newReader, err := zip.NewReader(reader, int64(len(data)))
	if err != nil {
		panic(err)
	}
	return &ZipEntry{abs, newReader}
}

func newDirEntry(path string) Entry {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{abs}
}

func newCompositeEntry(path string) CompositeEntry {
	var entries CompositeEntry
	split := strings.Split(path, pathSeparator)
	for _, s := range split {
		entry := newEntry(s)
		entries = append(entries, entry)
	}
	return entries
}

func newWildcardEntry(dirpath string) CompositeEntry {
	baseDir := dirpath[:len(dirpath)-1]
	//fmt.Println(baseDir)
	entry := CompositeEntry{}
	walkfunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, "zip") {
			entry = append(entry, newZipEntry(path))
		}
		return nil
	}
	err := filepath.Walk(baseDir, walkfunc)
	if err != nil {
		panic(err)
	}
	return entry
}
