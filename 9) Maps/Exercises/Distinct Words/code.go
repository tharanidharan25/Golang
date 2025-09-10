/*

Distinct Words
Complete the countDistinctWords function using a map. It should take a slice of strings and return the total count of distinct words across all the strings. Assume words are separated by spaces. Casing should not matter. (e.g., "Hello" and "hello" should be considered the same word).

For example:

messages := []string{"Hello world", "hello there", "General Kenobi"}
count := countDistinctWords(messages)

count should be 5 as the distinct words are "hello", "world", "there", "general" and "kenobi" irrespective of casing.

*/

package main

import (
	"fmt"
	"strings"
)

func countDistinctWords(messages []string) int {
	count := map[string]struct{}{}
	for _, sentence := range messages {
		for _, eachWord := range strings.Split(sentence, " ") {
			if _, ok := count[strings.ToLower(eachWord)]; !ok {
				count[strings.ToLower(eachWord)] = struct{}{}
			}
		}
	}
	return len(count)
}

func main() {
	distinctWords := countDistinctWords([]string{"WTS Arcanite Bar! Cheaper than AH", "Do you need an Arcanite Bar!"})
	fmt.Println("distinctWords:", distinctWords)
	distinctWords = countDistinctWords([]string{"Could you give me a number crunch real quick?", "Looks like we have a 32.33% (repeating of course) percentage of survival."})
	fmt.Println("distinctWords:", distinctWords)
	distinctWords = countDistinctWords([]string{
		"LF9M UBRS need all",
		"LF8M UBRS need all",
		"LF7M UBRS need all",
		"LF6M UBRS need tanks and heals",
		"LF5M UBRS need tanks and heals",
		"LF4M UBRS need tanks and heals",
		"LF3M UBRS need tanks and healer",
		"LF2M UBRS need tanks",
		"LF1M UBRS need tank",
		"Group is full thanks!",
	})
	fmt.Println("distinctWords:", distinctWords)
}