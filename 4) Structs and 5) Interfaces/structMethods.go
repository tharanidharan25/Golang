package main

import "fmt"

type rect struct {
	width  int
	height int
}

// area has a receiver of (r rect)
// rect is the struct
// r is the placeholder
func (r rect) area() int {
	return r.width * r.height
}

func main() {
	r := rect{
		width:  5,
		height: 10,
	}

	fmt.Println(r.area())
}