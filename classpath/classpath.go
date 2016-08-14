package classpath

import "os"
import "path/filepath"

// ref https://docs.oracle.com/javase/8/docs/technotes/tools/findingclasses.html
type Classpath struct {
	bootClasspath entry
	extClasspath  entry
	userClasspath entry
}

func Parse(userCpOption string) *Classpath {
	var cp = &Classpath{}

	var javaHome = os.Getenv("JAVA_HOME")
	cp.bootClasspath = newEntry(filepath.Join(javaHome, "jre", "lib", "*"))
	cp.extClasspath = newEntry(filepath.Join(javaHome, "jre", "lib", "ext", "*"))

	if userCpOption == "" {
		userCpOption = "."
	}
	cp.userClasspath = newEntry(userCpOption)

	return cp
}

func (self *Classpath) ReadClass(className string) ([]byte, error) {

	// TODO:
	return nil, nil
}
