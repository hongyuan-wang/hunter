package classpath

type Entry interface {
	toString() string
	readClass(className string) ([]byte, error)
}

func newEntry(path string) Entry {

	// TODO:
	return nil
}
