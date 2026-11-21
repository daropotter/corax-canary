package canary

import "strings"

// Greet returns a greeting for name, or for "world" when name is blank.
func Greet(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "world"
	}
	return "Hello, " + name + "!"
}
