package classpath

type compositeEntry []entry

func (self compositeEntry) toString() string {

	// TODO:
	return ""
}

func (self compositeEntry) readClass(className string) ([]byte, error) {

	return nil, nil
}

func newCompositeEntry(pathList string) compositeEntry {

	// TODO:
	return compositeEntry{}
}
