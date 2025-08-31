package main

import "fmt"

func printNumericValue(num interface{}) {
	// Type Switching with basic types in GO
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	case float64:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

// Example 2

func getExpenseReport(e expense) (to string, cost float64) {
	// Type switching with interfaces
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
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
	printNumericValue(1) // int
	printNumericValue("1") // string
	printNumericValue(1.00) // float64
	printNumericValue(true) // bool
	
	// Example 2
	e1 := email{
		isSubscribed: true,
		body: "Hello, there!",
		toAddress: "bob@example.com",
	}
	address, cost := getExpenseReport(e1)
	fmt.Println("address: ", address, "cost: ", cost)
	
	m1 := sms {
		isSubscribed: true,
		body: "Hello, there!",
		toPhoneNumber: "34275893",
	}
	phone, cost := getExpenseReport(m1)
	fmt.Println("phone: ", phone, "cost: ", cost)

}