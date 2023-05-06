package helper

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// asertt => Fail
// require => FailNow

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Gera")
	require.Equal(t, "Hell, Gera", result, "Result must be Hello, Gera") // harus sama 
	fmt.Println("TestHelloWorld with Require, Done")
}
func TestHelloWorldAssert1(t *testing.T) {
	result := HelloWorld("Gera")
	assert.Equal(t, "Hello, Gera", result, "Result must be Hello, Gera") // harus sama 
	fmt.Println("TestHelloWorld with Assert, Done")
}
func TestHelloWorldAssert2(t *testing.T) {
	result := HelloWorld("Gera")
	assert.NotEqual(t, "Hello, Anggara", result, "Result must be Hello, Gera") // harus beda
	fmt.Println("TestHelloWorld with Assert, Done")
}

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