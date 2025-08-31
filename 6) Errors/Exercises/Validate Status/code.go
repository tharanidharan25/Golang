package main

import (
	"errors"
	"fmt"
)

func validateStatus(status string) error {
	if len(status) == 0 {
		return errors.New("Status cannot be empty")
	} else if len(status) > 140 {
		return errors.New("Status is too lengthy, sinnadha eludhi palagu da")
	}
	return nil
}

func printRes(err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Status successfully set.")
	}
}

func main() {
	err := validateStatus("")
	printRes(err)
	
	err = validateStatus("fjhlksdajflsdjfoisdjfoisdjhfksdhnfkjsdhfksdhfksdhfsdfsdfsdkjfhksdhfksdhfkjsdhfkdshfkdshfklsadhfkjahfkjasdhfkjsdhfkjshadfkjhsdkfhsdkjfhskfhksdhffhsdkjf")
	printRes(err)
	
	err = validateStatus("Sleeping...😴💤")
	printRes(err)
}