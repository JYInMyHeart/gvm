package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry struct {
	readClass(className string)([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	} else if strings.Contains(path, "*") {
		return newWildcardEntry(path)
	} else if strings.HasSuffix(strings.ToLower(path), "jar") || strings.HasSuffix(strings.ToLower(path), "zip") {
		return newZipEntry(path)
	} else {
		return newDirEntry(path)
	}
}
