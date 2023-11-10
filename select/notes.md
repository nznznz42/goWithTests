# chapter 11

- select is a basically a match/switch statement designed for use with goroutines to allow them to wait on multiple communication operations
- a _select_ blocks until one of its cases can run and then it executes it, it chooses on at random if multiple are ready
- the "httptest" package is an easy way to spin up test servers to have controllable and reliable tests
- it uses the same interfaces as the actual http package which makes it easy to use both