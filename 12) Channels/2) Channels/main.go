package main

import (
	"fmt"
	"time"
)

type email struct {
	body string
	date time.Time
}

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

/*
Assignment
Our Mailio server isn't able to boot up until it receives the signal that its databases are all online, and it learns about them being online by waiting for tokens (empty structs) on a channel.

Run the code. It never exits! The channel passed to waitForDBs stays blocked, because it's only popping the first value off the channel.

Fix the waitForDBs function. It should pause execution until it receives a token for every database from the dbChan channel. Each time waitForDBs reads a token, the getDBsChannel goroutine will print a message to the console for you. The succinctly named numDBs input is the total number of databases.
*/

type testCase struct {
	numDBs int
}

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}

func main() {
	res := checkEmailAge([3]email{
		{
			body: "Music is a proud, temperamental mistress.",
			date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Have you heard of that website Boot.dev?",
			date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "It's awesome honestly.",
			date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	fmt.Println(res)

	res = checkEmailAge([3]email{
		{
			body: "I have stolen princesses back from sleeping barrow kings.",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "I burned down the town of Trebon",
			date: time.Date(2019, 6, 6, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "I have spent the night with Felurian and left with both my sanity and my life.",
			date: time.Date(2022, 7, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	fmt.Println(res)
	testCases := []testCase{
		{1},
		{3},
		{4},
	}
	passed, failed := 0, 0
	for _, test := range testCases {
		fmt.Printf(`---------------------------------`)
		fmt.Printf("\nTesting %v Databases...\n\n", test.numDBs)
		dbChan, count := getDBsChannel(test.numDBs)
		waitForDBs(test.numDBs, dbChan)
		for *count != test.numDBs {
			fmt.Println("...")
		}
		if len(dbChan) == 0 && *count == test.numDBs {
			passed++
			fmt.Printf(`
expected length: 0, count: %v
actual length:   %v, count: %v
PASS
`,
				test.numDBs, len(dbChan), *count)
		} else {
			failed++
			fmt.Printf(`
expected length: 0, count: %v
actual length:   %v, count: %v
FAIL
`,
				test.numDBs, len(dbChan), *count)
		}
	}
}
