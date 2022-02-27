package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Devri")
	if result != "Hello Devri" {
		panic("Result is not Hello Devri")
	}
}