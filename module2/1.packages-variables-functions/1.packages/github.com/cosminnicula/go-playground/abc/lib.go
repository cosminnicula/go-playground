// the directory inside a project is a Go package
// you cannot have multiple package names in a single directory; create multiple directories for multiple package names
package abc

// not accessible from outside the package
var privateVar = "private variable"

var PublicVar = "public variable"

// not accessible from outside the package
func privateFunction() string {
	return "private function"
}

func PublicFunction() string {
	return "public function"
}
