package classpath

// ref https://docs.oracle.com/javase/8/docs/technotes/tools/findingclasses.html
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(userCpOption string) *Classpath {
	var cp = &Classpath{}

	// TODO:
	return cp
}

func (self *Classpath) ReadClass(className string) ([]byte, error) {

	// TODO:
	return nil, nil
}
