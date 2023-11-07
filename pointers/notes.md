# chapter 6

- in go if a symbol starts with a lowercase symbol it is private to the package it is defined in
- go uses pass-by-value by default
- to pass in a reference to an object/variable/function etc you can use the "&thing" syntax similar to other languages
- you can use the * operator to declare a pointer as well as dereference it similar to how C does it
- struct pointers are automatically dereferenced
- errors in go are returned as values (good? bad? hard to say)
- the nil keyword in go is equivalent to null in other languages
- you can use errors.New("error message text") to create your own errors
- you can use t.Fatal to the kill the test if it fails
- errcheck is a cool linter for go which points out errors before build
- the compiler cannot help you with an unwanted null pointer because it results in a runtime exception