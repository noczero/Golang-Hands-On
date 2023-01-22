package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
t.Fail vs t.FailNow : when executed code is continoue vs code will stopped.
t.Error vs t.Fatal : using t.Fail + msg vs using T.FailNow + msg
t.Assert vs t.Require : from assert package hence more details : using t.Fail vs code using t.FailNow()

It's not recommend using if condition to check, just use assertion or require from assert.

Can skip using t.Skip

Use TestMain function to init and end unit test in one package.

Use table test for simply testing
*/

func TestTableHelloWorld(t *testing.T) {
	// table test

	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Satrya",
			request:  "Satrya",
			expected: "Hello Satrya",
		}, {
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		}, {
			name:     "Pratama",
			request:  "Pratama",
			expected: "Hello Pratama",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, HelloWorld(test.request))
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Satrya", func(t *testing.T) {
		// sub test
		result := HelloWorld("Satrya")
		assert.Equal(t, "Hello Satrya", result, "result should hello TRUE")
	})
	t.Run("Budi", func(t *testing.T) {
		// sub test
		result := HelloWorld("Budi")
		assert.Equal(t, "Hello Budi", result, "result should hello TRUE")
	})
}

// one package one TestMain function
func TestMain(m *testing.M) {
	//before
	fmt.Println("Before testing") //can do init database, etc
	m.Run()
	fmt.Println("End testing")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("TRUE")
	assert.Equal(t, "Hello TRUE", result, "result should hello TRUE")
}

func TestSkipFunction(t *testing.T) {
	t.Skip("Function test is skip")
	fmt.Println("Not Executed")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Satrya")
	if result != "Hello Satrya" {
		// error
		t.Fatal("Failed, not expected")
	}
}

func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("Budi")
	if result != "Hello Budi" {
		// error
		t.Fatal("Failed, not expected")
	}
}
