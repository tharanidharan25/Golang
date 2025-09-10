package main

import "fmt"

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func analyzeMessage(a *Analytics, msg Message) {
	if msg.Success {
		a.MessagesSucceeded += 1
	} else {
		// This will work too
		(*a).MessagesFailed += 1
	}
	a.MessagesTotal += 1
}

func main() {
	analytics := Analytics{
		MessagesTotal:     0,
		MessagesFailed:    0,
		MessagesSucceeded: 0,
	}
	analyzeMessage(&analytics, Message{
		Recipient: "mickey",
		Success:   true,
	})
	fmt.Println(analytics)
	analyzeMessage(&analytics, Message{Recipient: "minnie", Success: false})
	fmt.Println(analytics)
}