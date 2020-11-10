package classpath

import (
	"os"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	//used to fine and load the class file
	readClass(className string) ([]byte, Entry, error)
	//like toString method in java
	String() string
}

// func newEntry(path string) Entry {
// 	if strings.Contains(path, pathListSeparator) {
// 	}
// }
