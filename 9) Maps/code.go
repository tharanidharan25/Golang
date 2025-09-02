package main

import (
	"errors"
	"fmt"
)

/*
Assignment
It's important to keep up with privacy regulations and to respect our user's data. We need a function that will delete user records.

Complete the deleteIfNecessary function. The user struct has a scheduledForDeletion field that determines if they are scheduled for deletion or not.

If the user doesn't exist in the map, return the error not found.
If they exist but aren't scheduled for deletion, return deleted as false with no errors.
If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.
*/

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	for _, user := range users {
		if user.name == name && user.scheduledForDeletion {
			delete(users, name)
			return true, nil
		} else if user.name == name && !user.scheduledForDeletion {
			return false, nil
		}
	}
	return false, errors.New("User not found!")
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func runDeleteFunc() {
	deleted, err := deleteIfNecessary(map[string]user{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},"Eren")
	fmt.Println(deleted, err)
	
	deleted, err = deleteIfNecessary(map[string]user{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}}, "Erwin")
	fmt.Println(deleted, err)
}

/*

Assignment
Each time a user is sent a message, their username is logged in a slice. We want a more efficient way to count how many messages each user received.

Implement the updateCounts function. It takes as input:

messagedUsers: a slice of strings.
validUsers: a map of string -> int.
It should update the validUsers map with the number of times each user has received a message. Each string in the slice is a username, but they may not be valid. Only update the message count of valid users.

So, if "benji" is in the map and appears in the slice 3 times, the key "benji" in the map should have the value 3.

*/

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if _, ok := validUsers[user]; ok {
			validUsers[user]++
		}
	}
}

func runUpdateCounts() {
	newMap := map[string]int{"cersei": 0, "jaime": 0}
	updateCounts([]string{"cersei", "jaime", "cersei"}, newMap)
	fmt.Println(newMap)
	
	newMap = map[string]int{"cersei": 0, "jaime": 0, "tyrion": 0}
	updateCounts([]string{"cersei", "tyrion", "jaime", "tyrion", "tyrion"}, newMap)
	fmt.Println(newMap)
}

/*

Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.

The zero value of a map is nil.

We can create a map by using a literal or by using the make() function:

*/

func main() {
	// Declaring a map using make function
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 24
	ages["Mary"] = 21 // Overwrites 24

	// Initializing a map
	newAges := map[string]int{
		"John": 37,
		"Mary": 24,
		"Tharani": 22,
	}
	fmt.Println(ages, newAges)

	// len() function works on maps, it return the total number of key/value pairs
	fmt.Println(len(ages), len(newAges))

	// Insert an element
	ages["Tharani"] = 24

	// Get an Element
	fmt.Println("Tharani's age: ", ages["Tharani"])

	// Delete an Element
	delete(ages, "Tharani")
	fmt.Println(ages)

	// Checking if the key exists
	// ele will be the value if the key exists and ok will be true
	// If the key doesn't exist, ele will be the zero value of that value type and ok will be false
	ele, ok := ages["Tharani"]
	fmt.Println(ele, ok)
	ele, ok = ages["Mary"]
	fmt.Println(ele, ok)

	// Assignment
	runDeleteFunc()

	/* 
	Remember that you can check if a key is already present in a map by using the second return value from the index operation.
	You can combine an if statement with an assignment operation to use the variables inside the if block:
	*/
	names := map[string]int{}
	missingNames := []string{}

	if _, ok = names["Denna"]; !ok {
		missingNames = append(missingNames, "Denna")
	}
	fmt.Println(missingNames)
	// Assignment
	runUpdateCounts()
}