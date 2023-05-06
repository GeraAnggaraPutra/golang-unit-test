package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Gera")
	if result != "Hello, Gera" {
		// error
		panic("Result is not 'Hello, Gera'")
	}
}
func TestHelloWorldGera(t *testing.T) {
	result := HelloWorld("Gera")
	if result != "Hello, Gera" {
		// error
		panic("Result is not 'Hello, Gera'")
	}
}