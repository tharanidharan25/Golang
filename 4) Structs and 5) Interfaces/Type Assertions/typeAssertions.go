package main

import "fmt"

/* Type Assertion */
/*
	When working with interfaces in Go, every once-in-awhile you'll need access to the underlying type of an interface value. You can cast an interface to its underlying type using a type assertion.
*/
func getExpenseReport(e expense) (to string, cost float64) {
	// Asserting if e is of type email
	email, ok := e.(email)
	if ok {
		return email.toAddress, email.cost()
	}

	// Asserting if e is of type sms
	sms, ok := e.(sms)
	if ok {
		return sms.toPhoneNumber, sms.cost()
	}

	return "", 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func main() {
	newEmail := email{
		isSubscribed: true,
		body: "hello there",
		toAddress: "email@gmail.com",
	}
	to, cost := getExpenseReport(newEmail)
	fmt.Println("to: ", to, "cost:", cost)

	newSms := sms{
		isSubscribed:  false,
		body: "hello there",
		toPhoneNumber: "123456789",
	}
	to, cost = getExpenseReport(newSms)
	fmt.Println("to: ", to, "cost:", cost)
}