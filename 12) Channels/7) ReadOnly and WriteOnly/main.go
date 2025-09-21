/*
Tickers

time.Tick() - Standard library function that returns a channel that sends a value on a given interval.

time.After() - Sends a value once after the duration has passed.

time.Sleep() - Blocks the current goroutine for the specified duration of time.

These functions take time.Duration as an argument. For example:

	time.Tick(500 * time.Millisecond)

If we don't add 'time.Millisecond'(or another unit), it will default to nanoseconds. That's faster than we would want it to be.
*/

/*
Assignment
Like all good back-end engineers, we frequently save backup snapshots of the Mailio database.

Complete the saveBackups function.

It should read values from the snapshotTicker and saveAfter channels simultaneously and continuously.

If a value is received from snapshotTicker, call takeSnapshot()
If a value is received from saveAfter, call saveSnapshot() and return from the function: you're done.
If neither channel has a value ready, call waitForData() and then time.Sleep() for 500 milliseconds. After all, we want to show in the logs that the snapshot service is running.
*/
package main

// Assignment for Select in channels, Read-only channels and write-only channels
import (
	"fmt"
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot(logChan)
		case <-saveAfter:
			saveSnapshot(logChan)
			return
		default:
			waitForData(logChan)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}

func main() {
	snapShotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	logChan := make(chan string)
	go saveBackups(snapShotTicker, saveAfter, logChan)
	actualLogs := []string{}
	for actualLog := range logChan {
		fmt.Println(actualLog)
		actualLogs = append(actualLogs, actualLog)
	}
}
