package main

import "fmt"

func main() {
	msgLen := 10
	maxMsgLen := 20

	fmt.Println("Trying to send the message of length: ", msgLen, " and a max length of: ", maxMsgLen)

	// If-else Conditions

	if msgLen < maxMsgLen {
		fmt.Println("Message Sent")
	} else if msgLen == maxMsgLen {
		fmt.Println("Max limit reached and message sent")
	} else {
		fmt.Println("Message not sent")
	}

	// Initial Statement of If conditions
	// Limits the scope of the variable declared in the initial condition to this particular if-else condition

	if length := 10; length < maxMsgLen {
		fmt.Println("Message Sent")
	} else {
		fmt.Println("Message not sent")
	}

	// Switch Statements
	// "break" statement is not required in the 'switch' statements, it breaks by default
	// If you don't want it to break, use 'fallthrough' statement to make sure the program goes to the next case

	os := "windows"
	var creator string
	switch os {
		case "linux":
			creator = "Linus Torvalds"
		case "windows":
			creator = "Bill Gates"
		case "macOs":
			fallthrough
		case "Mac OS X":
			fallthrough
		case "mac":
			creator = "A Steve"
		
		default:
			creator = "Unknown"
	}

	fmt.Println(creator)
}