package main

import "fmt"

/*
We're hiring engineers at Textio, so time to brush up on the classic "Fizzbuzz" game, a coding exercise that has been dramatically overused in coding interviews across the world.

Complete the fizzbuzz function that prints the numbers 1 to 100 inclusive each on their own line, but replace multiples of 3 with the text fizz and multiples of 5 with buzz. Print fizzbuzz for multiples of 3 AND 5.
*/
func fizzbuzz(maxVal int) {
	// Go Supports standard && (and), || (or) and %(modulo) operators
	for i := 1; i <= maxVal; i++ {
		fmt.Print(i, ": ")
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println("Nothing")
		}
	}
}

func main() {
	// basic for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i, ", ")
	}

	// Omitting conditions from a for loop - To run infinite loop
	for i := 0; ; i++ {
		if i == 10 {
			fmt.Println("\nExceeded the maximum stack length: ", i)
			break
		}
	}

	// There is NO WHILE LOOP IN GO
	// While loop is just a for loop with only a condition
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("Plant is still growing... Current height is", plantHeight, "inches")
		plantHeight++
	}
	fmt.Println("Plant has grown now to ", plantHeight, "inches")

	fizzbuzz(10)
	// fizzbuzz(15)

	// Go also supports continue and break keywords
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}