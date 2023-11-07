# chapter 4

- arrays in go are declared using the syntax [n]T which is an array of n values of type T
- arrays are not resizable similar to other statically typed languages
- the "&v" format string is used to print the default format which is good for use with arrays
- the "range" keyword allows you to iterate over an array, on each iteration the index and value are returned either of which can be ignored using the _ identifier, this is similar to enumerate() in other languages
- go allows slices which allows for the existence of collections of indeterminate size
- the syntax for slices is basically the same as that of arrays except the size is omitted, i.e you would use []T to declare a slice of size T
- go does not allow the use of equality operators with slices
- both slices and arrays can be indexed using the arr[] syntax
- slices can be sliced themselves using the slice[low:high] syntax
- it is easy to see if any 2 variables in 2 slices are the same by using reflect.DeepEqual(x, y) 
- reflect.DeepEqual is not typesafe
- go has a coverage testing tool built in which can be used by running "go test -cover"
- go allows for variadic functions using the ... operator:
    ex: func example(nums ...[]int) []int {
        return nil
    }
- scopes can be created in go by enclosing statements in {}
