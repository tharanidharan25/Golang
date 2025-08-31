package main

import "fmt"

/*
	Many functions, especially those in the standard library can take an arbitrary number of final arguments. This is accomplished using "..." syntax in the function signature
*/

// Variadic function example:

// Spread Operator
// The spread operator allows us to pass a slice into a variadic function. The spread operator consists of three dots following the slice in the function call.
func concat(strs ...string) string {
	res := ""
	// strs are just a slice containing values of type string
	for i := 0; i < len(strs); i++ {
		res += strs[i]
	}
	return res
}

func main() {
	concatinated := concat("Hello ", "there, ", "friend!")
	// The familiar fmt.Println() and fmt.Sprintf() are variadic! fmt.Println() prints each element with space delimiters and a newline at the end.
	fmt.Println(concatinated)
}