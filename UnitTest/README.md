## Rules
1. file is ended with *_test*
2. function name is started with *Test*
3. has function paramater (t *testing.T) and doesn't have return value.

## How to run
1. Go on testing folder, then execute these command :
```bash
# run all testing on same folder
go test -v

# specific function name
go test -v -run=TestingHelloWorld2

# from root project
go test -v ./...

# just run subtest iniside test function
go test -v -run=TestSubTest/Satrya
```

## Table Test
Implements slice struct, subtest, and looping for dynamic parameter testing.

## Mock Test
Simulate the object data on repository layer and get testing on service layer

## Benchmark Test
Bencmark is used to calculate performance time of code execution.
The parameter is testing.B, the attribute for iteration is N where the code execute inside the iteration.
The result is how long in seconds the code been executed.

### Rules :
Name function starts with Benchmark word, no returm value.

### How to run :
```bash 
go test -v -run=NotMathUnitTest -bench=BenchmarkTest
```
