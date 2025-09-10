# Introduction to Pointers

As we have learned, a variable is a named location in memory that stores a value. We can manipulate the value of a variable by assigning a new value to it or by performing operations on it. When we assign a value to a variable, we are storing that value in a specific location in memory
```
x := 42
// "x" is the name of a location in memory. That location is storing the integer value of 42
```

# A Pointer Is a Variable
A pointer is a variable that stores the memory address of another variable. This means that a pointer "points to" the location of where the data is stored, not the actual data itself.

The * syntax defines a pointer:
```
var p *int
```

#### The & operator generates a pointer to its operand.
```
myString := "hello"
myStringPtr := &myString
```

# Dereference
The * operator dereferences a pointer to get the original value.
```
// set myString through the pointer
*myStringPtr = "world"

// read myString through the pointer
fmt.Printf("value of myString: %s\n", *myStringPtr)
// value of myString: world
```

# Pass by reference

Functions in Go generally pass variables by value, meaning that functions receive a copy of **most** non-composite types:

```
func increment(x int) {
    x++
    fmt.Println(x)
    // 6
}


func main() {
    x := 5
    increment(x)
    fmt.Println(x)
    // 5
}
```
The main function still prints 5 because the increment function received a copy of x.

One of the most common use cases for pointers in Go is to pass variables by reference, meaning that the function receives the address of the original variable, not a copy of the value. This allows the function to update the original variable's value.
```
func increment(x *int) {
    *x++
    fmt.Println(*x)
    // 6
}

func main() {
    x := 5
    increment(&x)
    fmt.Println(x)
    // 6
}
```

# Fields of Pointers
When your function receives a pointer to a struct, you might try to access a field like this and encounter an error:
```
msgTotal := *analytics.MessagesTotal
```

Instead, access it – like you'd normally do — using a selector expression.
```
msgTotal := analytics.MessagesTotal
```

This approach is the recommended, simplest way to access struct fields in Go, and is shorthand for:
```
(*analytics).MessagesTotal
```

# Nil Pointers
Pointers can be very dangerous.

If a pointer points to nothing (the zero value of the pointer type) then dereferencing it will cause a runtime error (a panic) that crashes the program.
> `Generally speaking, whenever you're dealing with pointers you should check if it's nil before trying to dereference it.`