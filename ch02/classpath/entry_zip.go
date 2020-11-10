package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

//ZipEntry stores the absolute path of ZIP or JAR files
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

//read file from zip
func (ze *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(ze.absPath)
	defer r.Close()
	if err != nil {
		return nil, nil, err
	}
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			defer rc.Close()
			if err != nil {
				return nil, nil, err
			}
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, ze, nil
		}
	}
	return nil, nil, errors.New("class not Found: " + className)
}

func (ze *ZipEntry) String() string {
	return ze.absPath
}
