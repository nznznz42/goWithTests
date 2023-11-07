# chapter 7

- maps in go are analagous to dictionaries in python or hashmaps in other languages
- they can be declared using the map[KeyType]ValueType syntax
- the key type must be a [comparable type](https://go.dev/ref/spec#Comparison_operators)
- maps can be modified without passing a pointer to them
- this works even though go has pass-by-value semantics because when a map is passed into a function, it is copied, but only the pointer to it, the data remains uncopied
- maps can be nil, they behave like empty maps while reading but will result in a runtime panic if you try to write to them, for this reason do not ever initialise an empty map
