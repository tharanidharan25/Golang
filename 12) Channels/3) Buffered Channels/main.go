/*
Assignment
We want to be able to send emails in batches. A writing goroutine will write an entire batch of email messages to a buffered channel, and later, once the channel is full, a reading goroutine will read all of the messages from the channel and send them out to our clients.

Complete the addEmailsToQueue function. It should create a buffered channel with a buffer large enough to store all of the emails it's given. It should then write the emails to the channel in order, and finally return the channel.
*/

package main

import "fmt"

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))
	for _, email := range emails {
		ch <- email
	}
	return ch
}

type testCase struct {
	emails []string
	expected int
}

func main() {
	tastCases := []testCase{
		{
			[]string{
				"To boldly go where no man has gone before.",
				"Live long and prosper.",
			},
			2,
		},
		{
			[]string{
				"The needs of the many outweigh the needs of the few, or the one.",
				"Change is the essential process of all existence.",
				"Resistance is futile.",
			},
			3,
		},
	}
	for _, test := range tastCases {
		ch := addEmailsToQueue(test.emails)
		fmt.Println(len(ch))
	}
}