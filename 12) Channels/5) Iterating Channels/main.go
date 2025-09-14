/*
Assignment
It's that time again, Mailio is hiring and we've been assigned to do the interview. The Fibonacci sequence is Mailio's interview problem of choice. We've been tasked with building a small toy program we can use in the interview.

Complete the concurrentFib function. It should:

Create a new channel of ints
Call fibonacci concurrently
Use a range loop to read from the channel and append the values to a slice
Return the slice
*/
package main

import (
	"fmt"
	"slices"
)

func concurrentFib(n int) []int {
	res := []int{}
	ch := make(chan int)
	go fibonacci(n, ch)
	for val := range ch {
		res = append(res, val)
	}
	return res
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	type testCase struct {
		n int
		expected []int
	}

	testCases := []testCase{
		{5, []int{0, 1, 1, 2, 3}},
		{3, []int{0, 1, 1}},
	}

	for idx, test := range testCases {
		res := concurrentFib(test.n)
		if !slices.Equal(res, test.expected) {
			fmt.Println("Testcase", idx, "Failed. Result -->", res)
		} else {
			fmt.Println("Testcase", idx, "Passed. Result -->", res)
		}
	}
}
