package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// asertt => Fail
// require => FailNow
type Tests struct {
	name, request, expected string
}

func TestTableHelloWorld(t *testing.T) {
	tests := []Tests{
		{
			name:     "Gera",
			request:  "Gera",
			expected: "Hello, Gera",
		},
		{
			name:     "Anggara",
			request:  "Anggara",
			expected: "Hello, Anggara",
		},
		{
			name:     "Putra",
			request:  "Putra",
			expected: "Hello, Putrsa",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be Hello, " + test.name)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Gera1", func(t *testing.T) {
		result := HelloWorld("Gera")
		require.Equal(t, "Hello, Gera", result, "Result must be Hello, Gera") // harus sama
	})
	t.Run("Gera2", func(t *testing.T) {
		result := HelloWorld("Gera2")
		require.Equal(t, "Hello, Gera2", result, "Result must be Hello, Gera2") // harus sama
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac Os")
	}

	result := HelloWorld("Gera")
	require.Equal(t, "Hello, Gera", result, "Result must be 'Hello, Gera'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Gera")
	require.Equal(t, "Hello, Gera", result, "Result must be Hello, Gera") // harus sama
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
