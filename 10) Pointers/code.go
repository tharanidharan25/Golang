package main

import (
	"fmt"
	"strings"
)

type Message struct {
	Recipient string
	Text      string
}

func getMessageText(m Message) string {
	// The & operator generates a pointer to its operand.
	return fmt.Sprintf(`
To: %v
Message: %v
`, &m.Recipient, &m.Text)
}

func removeProfanity(message *string) {
	*message = strings.ReplaceAll(*message, "fubb", "****")
	*message = strings.ReplaceAll(*message, "shizz", "****")
	*message = strings.ReplaceAll(*message, "witch", "*****")
}

func main() {
	fmt.Println(getMessageText(Message{
		"Tharani",
		"Hi Hello",
	}))
	ip := "English, motherfubber, do you speak it?"
	input := &ip
	removeProfanity(input)
	fmt.Println(ip)
}