# chapter 3

- go only has a single type of loop i.e for loops
- the syntax is basically C-like but does not use () to contain the index and range within the loop
- go has inbuilt support for benchmarks which can be accessed using "*testing.B" which gives us access to b.N which is the number of times the code will be executed to test its performance
- b.N is not user definable and is determined by the go runtime
- to run benchmarks use:
    1. on linux: "go test -bench=."
    2. on windows: "go test -bench=".""
- the benchmarks results will give you the value of b.N and time/op
- benchmarks are run sequentially 