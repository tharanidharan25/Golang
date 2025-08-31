package main

import "fmt"

func main() {
	// Basic variables (Refer .md)
	var smsSendingLimit int
	var costPerSms float64
	var hasPermission bool
	var username string

	fmt.Print("Basic variables")
	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSms, hasPermission, username)

	// Short Variable declaration using walrus operator (Refer .md)
	message := "Happy Birthday"
	age := 21
	fmt.Print("Short Variable declaration using walrus operator < := >: ")
	fmt.Println(message, age)

	// Same Line Declaration (Same as Python)
	average, display := .23, true
	fmt.Print("Same Line Declaration: ")
	fmt.Println(average, display)
	
	// Constant Variable (Same as JS)
	const pi float64 = 3.14
	fmt.Print("Constant Variable: ")
	fmt.Println(pi)

	// Formatting Strings in GO (Refer .md)
	// Learn more about "rune" in GO
	

}