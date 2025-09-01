/*
Message Tagger
Textio needs a way to tag messages based on specific criteria.

Assignment
Complete the tagMessages function. It should take a slice of sms messages, and a function (that takes a sms as input and returns a slice of strings) as inputs. And it should return a slice of sms messages.
It should loop through each message and set the tags to the result of the passed in function.
Be sure to modify the messages of the original slice using bracket notation messages[i].
Ensure the tags field contains an initialized slice. No nil slices.
Complete the tagger function. It should take an sms message and return a slice of strings.
For any message that contains "urgent" (regardless of casing) in the content, the Urgent tag should be applied first.
For any message that contains "sale" (regardless of casing), the Promo tag should be applied second.
Regardless of casing just means that even "uRGent" or "SALE" should trigger the tag.

Example usage:

messages := []sms{
	{id: "001", content: "Urgent! Last chance to see!"},
	{id: "002", content: "Big sale on all items!"},
	// Additional messages...
}
taggedMessages := tagMessages(messages, tagger)
// `taggedMessages` will now have tags based on the content.
// 001 = [Urgent]
// 002 = [Promo]

*/

package main

import (
	"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	res := []sms{}
	for i := 0; i < len(messages); i++ {
		tags := tagger(messages[i])
		if tags != nil {
			messages[i].tags = tags
			res = append(res, messages[i])
		}
	}
	return res
}

func tagger(msg sms) []string {
	tags := []string{}
	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}

func main() {
	fmt.Println(tagMessages([]sms{{id: "004", content: "Sale! Don't miss out on these urgent promotions!"}}, tagger))
	fmt.Println(tagMessages([]sms{{id: "005", content: "i nEEd URgEnt help, my FROZEN FLAME was used"}, {id: "006", content: "wAnt to saLE 200x heavy leather"}}, tagger))
}
