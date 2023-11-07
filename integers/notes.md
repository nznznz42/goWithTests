# chapter 2

- make sure there is only one package per directory
- format strings in go are C-like
- the walrus operator (:=) can be used to declare and initialise a variable at the same time similar to pinescript
- you can write (x, y int) instead of (x int, y int) in a function signature if the variables are of the same type
- the TDD workflow should go like: 
    1. Red (write a test that fails first)
    2. Green (write the minimal amount of code required to make tests pass)
    3. Refactor (refactor the code to retain functionality while being more robust)
- go has support for examples which are functions that are displayed as package docs and verified by running them as tests
