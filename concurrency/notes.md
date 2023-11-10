# chapter 10

- go supports concurrency through _goroutines_ (not identical to kotlin coroutines but used similarly)
- the syntax to invoke one is
```
    go func_name(args)
```
- goroutines run in the same address space so access to shared memory must be synchronised. the "sync" package provides useful primitives to achieve this.
- channels help organise and control communications between the different threads and help avoid race conditions.
- the syntax to use channels is:
```
    ch := make(chan int) // create channel
    ch <- v              // Send v to channel ch.
    v := <-ch            // Receive from ch, and
                         // assign value to v.
```
- go comes with a race detector as part of the tool chain, to use it run:
```
    go run -race package.go
```

