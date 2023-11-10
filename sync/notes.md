# chapter 13

- the sync package provides basic synchronization primitives such as mutexes, WaitGroup and Once types
- while channels allow you to accomplish the same thing at higher levels, the synch package is meant more for lower level multithreading tasks
- in general, channels are useful when the data being passed changes ownership and mutexes are useful when the data is to be managed but does not change owners
- go vet is a handy tool to test for subtle bugs
- avoid embedding mutexes into a struct to keep things private
- a mutex must not be copied after first use
- a WaitGroup waits for a collection of goroutines to finish and blocks until they are