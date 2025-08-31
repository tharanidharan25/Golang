package main

import "fmt"

/*
Password Strength:
	As part of improving security, Textio wants to enforce a new password policy. A valid password must meet the following criteria:

	At least 5 characters long but no more than 12 characters.
	Contains at least one uppercase letter.
	Contains at least one digit.

Assignment:
	Implement the isValidPassword function by looping through each character in the password string. Make sure the password is long enough and includes at least one uppercase letter and one digit.

	Assume passwords consist of ASCII characters only.
*/

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	containsDigit := false
	containsUpperCase := false
	containsLetters := false
	for i := 0; i < len(password); i++ {
		c := password[i]
		if c >= '0' && c <= '9' {
			containsDigit = true
		} else if c >= 'A' && c <= 'Z' {
			containsUpperCase =  true
		} else {
			containsLetters = true
		}
	}
	return containsDigit && containsUpperCase && containsLetters
}

func main() {
	fmt.Println(isValidPassword("Pass123"))
	fmt.Println(isValidPassword("pas"))
	fmt.Println(isValidPassword("Password"))
	fmt.Println(isValidPassword("123456"))
	fmt.Println(isValidPassword("VeryLongPassword1"))
	fmt.Println(isValidPassword("Shor"))
}