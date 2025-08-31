package main

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs %v to be sent to %v can not be sent", cost, recipient)
}

func main() {
	errorMsg := getSMSErrorString(100, "Tharani")
	fmt.Println(errorMsg)
}