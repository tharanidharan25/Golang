/*
Connections
	Textio has group chats that make communicating with multiple people much more efficient--if the chat doesn't descend into chaos. Instead of sending the message multiple times to individual people, you send one message to all of them at once.

Assignment
	Complete the countConnections function that takes an integer groupSize representing the number of people in the group chat and returns an integer representing the number of connections between them. For each additional person in the group, the number of new connections continues to grow. Use a for loop to accumulate the number of connections instead of directly using a mathematical formula.

	To make sure you’re picturing it right:

	If there are two people, how many possible connections exist between them?
	If you add a third person, how many new connections are created with the rest?
*/

package main

import "fmt"

func countConnections(groupSize int) int {
	res := 0
	for i := 1; i <= groupSize; i++ {
		res += i - 1
	}
	return res
}

func main() {
	res := countConnections(5)
	fmt.Println("5: ", res)
}