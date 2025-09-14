/*
Assignment
At Mailio we're all about keeping track of what our systems are up to with great logging and telemetry.

The sendReports function sends out a batch of reports to our clients and reports back how many were sent across a channel. It closes the channel when it's done.

Complete the countReports function. It should:

Use an infinite for loop to read from the channel:
If the channel is closed, break out of the loop
Otherwise, keep a running total of the number of reports sent
Return the total number of reports sent
*/

package main

import "fmt"

func countReports(numSentCh chan int) int {
	total := 0
	for i := 0; i < 1; {
		val, ok := <-numSentCh
		if ok == false {
			break
		}
		total += val
	}
	return total
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

type testCase struct {
	numBatches int
	expected int
}

func main() {
	testCases := []testCase{
		{3, 114}, {4, 198},
		{0, 0}, {1, 15}, {6, 435},
	}
	for idx, test := range testCases {
		numSentCh := make(chan int)
		go sendReports(test.numBatches, numSentCh)
		output := countReports(numSentCh)
		if output == test.expected {
			fmt.Println("Testcase", idx, "Passed.", "Result: ", output)
		} else {
			fmt.Println("Testcase", idx, "Failed.", "Result: ", output)
		}
	}
}
