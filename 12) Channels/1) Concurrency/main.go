package main

import (
	"fmt"
	"time"
)

/*
Assignment
At Mailio we send a lot of network requests. Each email we send must go out over the internet. To serve our millions of customers, we need a single Go program to be capable of sending thousands of emails at once.

Edit the sendEmail() function to execute its anonymous function concurrently so that the "received" message prints after the "sent" message.
*/
func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func main() {
	test("Hello there Kaladin!")
	test("Hi there Shallan!")
	test("Hey there Dalinar!")
}
