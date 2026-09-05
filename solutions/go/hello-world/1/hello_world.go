package hello

const testVersion = 2

// It's good style to write a comment here documenting HelloWorld.
func HelloWorld(s string) string {
	if s == "" {
		return "Hello, World!"
	}

	return "Hello, " + s + "!"
}