package classpath

import "path/filepath"

type dirEntry struct {
	absolutePath string
}

func (self dirEntry) toString() string {
	return self.absolutePath
}

func (self dirEntry) readClass(className string) ([]byte, error) {

	// TODO:
	return nil, nil
}

func newDirEntry(path string) dirEntry {
	var absolutePath, _ = filepath.Abs(path)
	return dirEntry{absolutePath: absolutePath}
}
