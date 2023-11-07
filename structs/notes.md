# chapter 5

- structs in go are equivalent to C and rust structs in use and function, i.e they're used to create custom types which consist of named fields.
- the syntax to create a struct is:
```
    type NAME struct {
        field1 type
        field2 type
            .
            .
            .
        fieldn type
    }
```
- struct fields can be accessed using the . operator much like any other language
- methods can be defined on and bound to structs similar to how methods can be defined in classes in OOP languages
- the syntax for declaring methods on a struct in go is:
    func (recName RecType) methodName(args)
- the convention in go is to have the receiver variable be the first letter of the reciever type
- interfaces are a type in go and are somewhat different to interfaces in other languages
- the empty interface can be used by code that handles values of unknown types
- other than this interfaces in go are used as usual, i.e they're used to define a contract which every implementor MUST adhere to
- interfaces are used in go to allow parameteric polymorphism (was this how they did it before adding generics???)
- the format string "%g" can be used in place of "%f" when we need more precision
