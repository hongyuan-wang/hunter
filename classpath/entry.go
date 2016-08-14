package classpath

import (
	"os"
	"strings"
)

type entry interface {
	toString() string
	readClass(className string) ([]byte, error)
}

func newEntry(path string) entry {

	if strings.Contains(path, string(os.PathListSeparator)) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, "JAR") ||
		strings.HasSuffix(path, "zip") || strings.HasSuffix(path, "ZIP") {

		return newJarEntry(path)

	}

	return newDirEntry(path)
}
