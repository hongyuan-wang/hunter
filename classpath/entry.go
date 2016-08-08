package classpath

type Entry interface {
	String() string
	readClass(className string) ([]byte, Entry, error)
}
