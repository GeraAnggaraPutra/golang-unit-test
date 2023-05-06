package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldGera(t *testing.T) {
	result := HelloWorld("Gera")
	if result != "Hello, Gera" {
		// error
		t.Error("Result harusnya 'Hello, Gera'")
	}
	fmt.Println("TestHelloWorldGera, Done")
}
func TestHelloWorldAnggara(t *testing.T) {
	result := HelloWorld("Anggara")
	if result != "Hello, Anggara" {
		// error
		t.Fatal("Result harusnya 'Hello, Anggara'")
	}
	fmt.Println("TestHelloWorldAnggara, Done")
}