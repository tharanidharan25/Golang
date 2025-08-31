package main

import (
	"fmt"
)

/*
	Structs in Go are often used to represent data that you might use a dictionary or object for in other languages.
*/

// It's conventional to declare structs at the package level (outside of any function)
// so they can be reused.

// The 'wheel' struct must be declared before 'car' because 'car' uses it. This applies only inside a function, otherwise it can be in any order

type car struct {
	brand      string
	model      string
	doors      int
	mileage    int
	frontWheel wheel
	backWheel  wheel
}

type wheel struct {
	radius   int
	material string
}

// Anonymous structs inside normal structs
type person struct {
	name string
	age int
	address struct {
		street string
		city string
		state string
		zipCode string
	}
}

// EMBEDDED STRUCTS
/*
	Go is not an object oriented language. However, embedded structs provide a kind of data-only inheritance that can be useful at times. Keep in mind, Go doesn't support classes or inheritance in the complete sense, but embedded structs are a way to elevate and share fields between struct definitions.
*/
type truck struct {
	car
	bedSize int
}

// EMPTY STRUCT
// named empty struct
type emptyStruct struct {}

func main() {
	// Defining a struct
	myCar := car{}
	// Accessing the fields of a struct
	myCar.brand = "Toyota"
	myCar.frontWheel.radius = 5
	fmt.Println(myCar.brand)

	// Anonymous Structs in GO.
	// An anonymous struct is just like a normal struct, but it is defined without a name and therefore cannot be referenced elsewhere in the code.
	// If a struct is only meant to be used once, then it makes sense to declare it in such a way(anonymous structs) that developers down the road won’t be tempted to accidentally use it again.
	mySecondCar := struct {
		brand string
		model string
	} {
		brand: "Toyota",
		model: "Camry",
	}

	fmt.Println(mySecondCar)
	fmt.Println(mySecondCar.brand)

	me := person {
		name: "Tharani",
		age: 24,
		address: struct {
			street string
			city string
			state string
			zipCode string
		} {
			street: "123 Main St",
			city: "Anytown",
			state: "CA",
			zipCode: "12345",
		},
	}

	fmt.Println(me)

	// EMBEDDED STRUCTS
	lanesTruck := truck {
		bedSize: 10,
		car: car {
			brand: "Toyota",
			model: "Tundra",
		},
	}
	/*
		Embedded vs Nested:
		unlike nested structs, embedded struct's fields are accessed at the top level like normal fields
	*/
	fmt.Println(lanesTruck)
	fmt.Println(lanesTruck.brand)
	fmt.Println(lanesTruck.model)

	// EMPTY STRUCTS
	// Cool thing about empty structs in that they're the smalles possible type in Go, they take up zero bytes of memory.
	// Empty structs are used in GO as a unary value.
	// Anonymous empty structs
	empty := struct{}{}
	fmt.Println(empty)

	// named empty Struct
	namedEmpty := emptyStruct{}
	fmt.Println(namedEmpty)
}