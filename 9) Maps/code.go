package main

/*

Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.

The zero value of a map is nil.

We can create a map by using a literal or by using the make() function:

*/

func main() {
	// Declaring a map using make function
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 24
	ages["Mary"] = 21 // Overwrites 24

	// Initializing a map
	ages = map[string]int{
		"John": 37,
		"Mary": 24,
		"Tharani": 22,
	}
}