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
