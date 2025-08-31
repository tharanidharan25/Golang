package main

import (
	"errors"
	"fmt"
	"log"
)

// Errors package provided by GO's standard library

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	cost, error := sendSMS(msgToCustomer)
	if error != nil {
		return 0, error
	}
	spouseMsgCost, error := sendSMS(msgToSpouse)
	if error != nil {
		return 0, error
	}
	return cost + spouseMsgCost, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 10
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

// Example 2
/*
	Errors are just interfaces, we can build our own custom types that implement the 'error' interface.

Assignment:
	Our users frequently try to run custom analytics queries on their message deliverability metrics, and end up writing a bad query that tries to divide a number by zero. It's become such a problem that we need to make a new type of error for division by zero.

	Update the code so that the divideError type implements the error interface. Its Error() method should just return a string formatted in the following way:

	can not divide DIVIDEND by zero

	Where DIVIDEND is the actual dividend of the divideError. Use the %v verb to format the dividend as a float.
*/

type divideError struct {
	dividend float64
}

func (d divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", d.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

// Example 3
// Go's standard library "errors" example
func divideNewError(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by Zero")
	}
	return x / y, nil
}


func main() {
	totalCost, error := sendSMSToCouple("dkjfafs", "sdjfsfsf")
	if error != nil {
		fmt.Println("Error encountered: ", error)
		// return
	} else {
		fmt.Println("Msg sent successfully. Total Cost = ", totalCost)
	}

	res, error := divide(100, 50)
	if error != nil {
		fmt.Println("Error encountered: ", error)
	} else {
		fmt.Println("Result: ", res)
	}

	res, error = divide(100, 0)
	if error != nil {
		fmt.Println("Error encountered: ", error)
	} else {
		fmt.Println("Result: ", res)
	}

	// errors.New() is a function provided by GO standard library errors, to convert strings to error type
	res, err := divideNewError(2.0, 0.0)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("Result: ", res)
	}

	res, err = divideNewError(20.0, 2.0)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("Result: ", res)
	}

	// Learn more about panic/recover - Do not use this in normal cases

	// For a truly unrecoverable error, to cleanly exit the program, use log.fatal() to print the error and exit the program.
	log.Fatal("Unrecoverable Error encountered... Exiting the program")
}
