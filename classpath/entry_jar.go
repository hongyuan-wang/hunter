package classpath

import "path/filepath"

type jarEntry struct {
	absolutePath string
}

func (self jarEntry) toString() string {
	return self.absolutePath
}

func (self jarEntry) readClass(className string) ([]byte, error) {

	// TODO:
	return nil, nil
}

func newJarEntry(path string) jarEntry {

	var absolutePath, _ = filepath.Abs(path)
	return jarEntry{absolutePath: absolutePath}
}
