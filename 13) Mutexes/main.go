package main

import (
	"fmt"
	"sync"
	"time"
)

type testCase struct {
	email string
	count int
}

/*
Assignment
We send emails across many different goroutines at Mailio. To keep track of how many we've sent to a given email address, we use an in-memory map.

Our safeCounter struct is unsafe! Update the inc() and val() methods so that they utilize the safeCounter's mutex to ensure that the map is not accessed by multiple goroutines at the same time.
*/

type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}

func (sc *safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc *safeCounter) val(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.slowVal(key)
}

// don't touch below this line

func (sc *safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc *safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}

// Maps are NOT thread-safe for concurrent use. If we have multiple goroutines accessing the same map, and at least one of them is writing to the map, we must lock the maps with a mutex.
func main() {
	passCount := 0
	failCount := 0
	testCases := []testCase{
		{"ken@waystar.com", 23},
		{"roman@waystar.com", 67},
		{"logan@waystar.com", 31},
		{"shiv@waystar.com", 453},
	}
	for _, test := range testCases {
		sc := safeCounter{
			counts: make(map[string]int),
			mu: &sync.Mutex{},
		}
		var wg sync.WaitGroup
		for i := 0; i < test.count; i++ {
			wg.Add(1)
			go func(email string) {
				sc.inc(email)
				wg.Done()
			}(test.email)
		}
		wg.Wait()
		if output := sc.val(test.email); output != test.count {
			failCount++
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
email: %v
count: %v
expected count: %v
actual count:   %v
`, test.email, test.count, test.count, output)
		}
	}
	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
