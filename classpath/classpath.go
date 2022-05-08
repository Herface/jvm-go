package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

type Classpath struct {
	// 启动类路径
	bootPath Entry
	// 扩展类路径
	extPath Entry
	// 用户类路径
	userPath Entry
}

func (this *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = strings.ReplaceAll(className, ".", "/") + ".class"
	if class, entry, err := this.bootPath.ReadClass(className); err == nil {
		return class, entry, err
	}
	if class, entry, err := this.extPath.ReadClass(className); err == nil {
		return class, entry, err
	}
	return this.userPath.ReadClass(className)

}
func (this *Classpath) String() string {
	return this.bootPath.String() + pathSeparator + this.extPath.String() + pathSeparator + this.userPath.String()
}
func (this *Classpath) parseBootAndExtPath(jrePath string) {
	jrePath = getJrePath(jrePath)
	libPath := filepath.Join(jrePath, "lib", "*")
	this.bootPath = newWildcardEntry(libPath)
	extPath := filepath.Join(jrePath, "lib", "ext", "*")
	this.extPath = newWildcardEntry(extPath)
}

func (this *Classpath) parseUserPath(userPath string) {
	if userPath == "" {
		userPath = "."
	}
	this.userPath = newDirEntry(userPath)
}

// 获取 jre类库路径
func getJrePath(jrePath string) string {
	if jrePath != "" && exist(jrePath) {
		return jrePath
	} else if javaHome := os.Getenv("JAVA_HOME"); javaHome != "" {
		return filepath.Join(javaHome, "jre")
	}
	panic("could not find jre path")
}

func exist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true

}
func NewClassPath(jrePath, userPath string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtPath(jrePath)
	cp.parseUserPath(userPath)
	return cp
}
