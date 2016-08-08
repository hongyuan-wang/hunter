package classpath

// ref https://docs.oracle.com/javase/8/docs/technotes/tools/findingclasses.html
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}
