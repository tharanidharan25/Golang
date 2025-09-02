/*

Any type can be used as the value in a map, but keys are more restrictive.

Read the following section of the official Go blog - https://go.dev/blog/maps:

As mentioned earlier, map keys may be of any type that is comparable. The language spec defines this precisely, but in short, comparable types are boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types. Notably absent from the list are slices, maps, and functions; these types cannot be compared using ==, and may not be used as map keys.

It's obvious that strings, ints, and other basic types should be available as map keys, but perhaps unexpected are struct keys. Struct can be used to key data by multiple dimensions.

*/

package main

import "fmt"

type Key struct {
    Path, Country string
}

func main() {
	hits := make(map[Key]int)

	// When an Indian visits the home page, incrementing (and possibly creating) the appropriate counter is a one-liner:
	hits[Key{"/", "in"}]++
	fmt.Println(hits)
	
	// And it’s similarly straightforward to see how many Swiss people have read the spec:
	n := hits[Key{"/ref/spec", "ch"}]
	fmt.Println(n)
}