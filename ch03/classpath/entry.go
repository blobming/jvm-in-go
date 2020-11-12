package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

// Entry contains readClass method and String method -> interface
type Entry interface {
	//used to fine and load the class file
	readClass(className string) ([]byte, Entry, error)
	//like toString method in java
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
