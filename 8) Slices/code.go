package main

import (
	"fmt"
)

func main() {
	// Declaring an array
	// Integer type array of size 10
	var myInts [10]int
	fmt.Println(myInts)
	// String type of size 5
	var myStrs [5]string
	fmt.Print(myStrs)
	fmt.Printf("%T\n", myStrs[0])

	// Initializing an array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	/*
		99 times out of 100, we'll use a slice instead of an array when working with ordered lists.

		Arrays are fixed in size. Once we create an array like [10]int, we can't add a 11th element.
		
		The zero value of an array or slice in nil
		var []myInt - This is equal to nil
	*/

	// Non-nil slices always have an underlying array, though it isn't always specified explicitly. To explicitly create a slice on top of an array we can do:
	/*
		arrayname[lowIndex:highIndex]
		arrayname[lowIndex:]
		arrayname[:highIndex]
		arrayname[:]
		
		Both can be omitted to use the entire array on that side of the colon.

		Note: lowIndex is inclusive and highIndex in exclusive
	*/
	mySlice := primes[1: 4]
	fmt.Println(mySlice)

	// Creating a slice without an explicit array - Remember, A slice will always be referenced to an array internally
	// We need "make" function to create a slice
	// make function has three arguments, make(<TYPE>, <LENGTH>, <CAPACITY>)
	// Capacity will be omitted most of the time, it will default to length.
	mySlice = make([]int, 5)
	fmt.Println(mySlice)
	
	mySlice = make([]int, 5, 10)
	fmt.Println(mySlice)

	// Slice created with make function will default to it's zero value

	// Declaring an array with values using make function
	myStrSlice := []string{"I", "love", "Go"}
	fmt.Println(myStrSlice)

	// Length of a slice is simply the number of elements it contains. It can be accessed using len() function
	fmt.Println("Length of slice: ", len(mySlice))

	// Capacity of a slice is the number of elements in the underlying array, Counting from first element in the slice. It can accessed using built-in cap() function
	fmt.Println("Capacity of slice: ", cap(mySlice))

	// indexing
	fmt.Println(myStrSlice[1])
}