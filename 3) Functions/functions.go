package main

import (
	"errors"
	"fmt"
)

/*
func - keyword
concat - function name
s1 string, s2 string - function parameters with their types
string - return type
*/
func concat(s1 string, s2 string) string {
	return s1 + s2
}

/* 
Multiple function parameters with same type can be seperated using comma and have type only for their last parameter of same type. (It should be sequential)
 */

func add(num1, num2 int) int {
	return num1 + num2
}

/* 
Ignoring the return value

We can use "_" identifier to ignore the return value

In GO, if we declare a variable and not use it, it'll throw an error, to get around it we can use blank identifier - "_".
Example below
 */

func getPoint() (x int, y int) {
	/*
		Note, I have return values with variable names in the function declaration, so when we return values, it'll be returned in that order. x will correspond to first value in the return statement, y will correspond to the second value in the return statement and so on. This is called Named return values
	*/
	return 3, 4
}

/* Named return values */
func calculator(a, b int) (mul int, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("Cannot Divide by Zero")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}

/* naked return is when we just write "return" statement and nothing else. GO will pickup the variable with the name in the named returns in the function signature and return it automatically. This should only be used in short functions
example below
*/

func nakedReturn() (x, y int) {
	x = 3
	y = 4
	return
}

/* 
	Functions as values - Functions are also values in GO, just like any other variable (Similar to JS). They can also be passed as a parameter to another function. Functions are just another type like int, float, bool and so on.
 */

func aggregator(a, b, c int, arithmetic func(int, int) int) (final int) {
	firstValue := arithmetic(a, b)
	final = arithmetic(firstValue, c)
	return
}

/* 
	"defer" keyword is a unique feature in GO. It lets the user to run any function just before the current functions returns. defer keyword is used to cleanup resources, close db connections, in file handlers and like.
	Example below
*/

func cleanUp() {
	fmt.Println("Cleaning up...")
}

func showDefer() (string) {
	fmt.Println("Preparing Clean Up...")
	defer cleanUp()
	a := "Cleaned Up."
	return a
}

/* GO also has closures similar to JS. Example Below*/
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func main() {
	fmt.Println(concat("Hello ", "World"))
	fmt.Println(add(100, 20))

	/* Return values declaration in function declaration and blank identifier usecase */
	x, _ := getPoint()
	/* Note that blank identifier is not throwing an error */
	fmt.Println(x)

	mul, div, err := calculator(10, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mul, div)

	a, b := nakedReturn()
	fmt.Println(a, b)

	result := aggregator(1, 2, 3, add)
	fmt.Println(result)

	/* GO also has anonymous functions similar to JS. Example below */
	mulResult := aggregator(5, 5, 5, func(a, b int) int {
		return a * b
	})
	fmt.Println(mulResult)

	deferRes := showDefer()
	fmt.Println(deferRes)

	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
}