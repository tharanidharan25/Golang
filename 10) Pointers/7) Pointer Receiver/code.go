package main

import "fmt"

type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

type nonPointerCar struct {
	color string
}

func (c nonPointerCar) setColor(color string) {
	c.color = color
}

// Assignment
/* Fix the bug in the code so that setMessage sets the message field of the given email struct, and the new value persists outside the scope of the setMessage method. */
type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

func main() {
	c := car{
		color: "White",
	}
	/* Methods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value. */
	// Notice 'c' is not a pointer in the calling function, but the method still gains access to a pointer to 'c'
	c.setColor("Green")
	fmt.Println(c.color) // Prints Green

	npCar := nonPointerCar {
		color: "White",
	}
	npCar.setColor("Green")
	fmt.Println(npCar.color) // Prints White, because of non-pointer receiver

	// Assignment
	e := email {
		message: "Edhu Nagarjuna vaaa",
		fromAddress: "Rajini",
		toAddress: "Us",
	}
	e.setMessage("Edhu Nagarjuna vaa!!!!!!😱😱")
	fmt.Println(e.message)
}