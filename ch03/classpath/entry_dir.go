package classpath

import (
	"io/ioutil"
	"path/filepath"
)

//DirEntry used to store the absolute path
type DirEntry struct {
	absDir string
}

//newDirEntry converts the parameter to abs path
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

//readClass returns the content of classfile
func (de *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(de.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, de, err
}

//String a simple method to return the absDir
func (de *DirEntry) String() string {
	return de.absDir
}
