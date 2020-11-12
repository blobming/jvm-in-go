package classpath

import (
	"os"
	"path/filepath"
)

//Classpath contains bootClasspath, extClasspath, userClasspath
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

//Parse parse boot and ext class path with jreOption
//Parse parse user class path with cpOption
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp

}

//ReadClass used to search and load the class file from bootClasspath, extClasspath, userClasspath
func (cp *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := cp.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := cp.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return cp.userClasspath.readClass(className)
}
func (cp *Classpath) String() string {
	return cp.userClasspath.String()
}

func (cp *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	cp.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	cp.extClasspath = newWildcardEntry(jreExtPath)
}

// if user does not provide the -classpath/-cp option, we use the current path as the user path
func (cp *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	cp.userClasspath = newEntry(cpOption)
}

//high-prioritize the user defined path or search from $Java_Home
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

// check if the path exists or not
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
