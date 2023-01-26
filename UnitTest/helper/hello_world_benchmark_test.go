package helper

import "testing"

/*
Bencmark is used to calculate performance time of code execution.
The parameter is testing.B, the attribute for iteration is N where the code execute inside the iteration.
The result is how long in seconds the code been executed.

Rules :
Name function starts with Benchmark word, no returm value.

How to run :
go test -v -run=NotMathUnitTest -bench=BenchmarkTest
*/

// Simple Benchmark
func BenchmarkHelloWorld(b *testing.B) {
	// iteration
	for i := 0; i < b.N; i++ {
		// invoke the function
		HelloWorld("Satrya")
	}
}

// Sub Benchmark
func BenchmarkHelloWorldSubBenchmark(b *testing.B) {
	b.Run("Bencmark1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Satrya PP")
		}
	})

	b.Run("Bencmark2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Satrya CC")
		}
	})
}

// Benchmark Table
func BenchmarkHelloWorldUsingTable(b *testing.B) {
	bencmarksTable := []struct {
		name, request string
	}{
		{
			name:    "HelloWorld(SatryaBudiPratama)",
			request: "Satrya",
		},
		{
			name:    "HelloWorld(SatryaBudiPratama222)",
			request: "Satrya222",
		},
	}

	for _, bencmark := range bencmarksTable {
		b.Run(bencmark.request, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bencmark.name)
			}
		})
	}
}
