/*
	- Implement the formatter interface with a method format that returns a formatted string.
	- Define structs that satisfy the formatter interface: plainText, bold, code.
	- The structs must all have a message field of type string
	plainText should return the message as is.
	- bold should wrap the message in two asterisks (**) to simulate bold text (e.g., **message**).
	- code should wrap the message in a single backtick (`) to simulate code block (e.g., `message`)
*/

package main

import "fmt"

type formatter interface {
	format() string
}

func sendMessage (msg formatter) string {
	return msg.format()
}

type plainText struct {
	message string
}

func (p plainText) format() string {
	return p.message
}

type bold struct {
	message string
}

func (b bold) format() string {
	return "**" + b.message + "**"
}

type code struct {
	message string
}

func (c code) format() string {
	return "`" + c.message + "`"
}

func main() {
	newText := plainText{
		message: "Hi, Hello, How are you?",
	}
	formattedText := sendMessage(newText)
	fmt.Println("Formatted Plain Text: ", formattedText)

	boldText := bold{
		message: "Hi, Hello, How are you?",
	}
	formattedText = sendMessage(boldText)
	fmt.Println("Formatted Bold Text: ", formattedText)
	
	codeText := code{
		message: "Hi, Hello, How are you?",
	}
	formattedText = sendMessage(codeText)
	fmt.Println("Formatted Code: ", formattedText)
}
