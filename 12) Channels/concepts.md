# Concurrency

#### What Is Concurrency?
Concurrency is the ability to perform multiple tasks at the same time. Typically, our code is executed one line at a time, one after the other. This is called sequential execution or synchronous execution.

If the computer we're running our code on has multiple cores, we can even execute multiple tasks at exactly the same time. If we're running on a single core, a single core executes code at almost the same time by switching between tasks very quickly. Either way, the code we write looks the same in Go and takes advantage of whatever resources are available.

#### How Does Concurrency Work in Go?
Go was designed to be concurrent, which is a trait fairly unique to Go. It excels at performing many tasks simultaneously safely using a simple syntax.

Concurrency is as simple as using the go keyword when calling a function:
```
go doSomething()
```
In the example above, doSomething() will be executed concurrently with the rest of the code in the function. The go keyword is used to spawn a new [goroutine](https://gobyexample.com/goroutines).

# Channels

#### Channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other.

### Creating a channel
Like maps and slices, channels must be created before use. They also use the same make keyword:
```
ch := make(chan int)
```

### Send data to a channel
```
ch <- 69
```
The `<-` operator is called the `channel operator`. Data flows in the direction of the arrow. `This operation will block until another goroutine is ready to receive the value`.

### Receive Data from a channel
```
v := <-ch
```
`This reads and removes a value from the channel and saves it into the variable v. This operation will block until there is a value in the channel to be read.`

### Reference Type
Like maps and slices, channels are reference types, meaning they are passed by reference by default

```
func send(ch chan int) {
    ch <- 99
}

func main() {
    ch := make(chan int, 1)
    send(ch)
    fmt.Println(<-ch) // Prints 99
}
```

### Blocking and Deadlocks
A [deadlock](https://yourbasic.org/golang/detect-deadlock/#:~:text=yourbasic.org%2Fgolang,look%20at%20this%20simple%20example.) is when a group of goroutines are all blocking so none of them can continue. This is a common bug that you need to watch out for in concurrent programming.

#### Receiving values from channel:
```
v := <-ch
```

However, sometimes we don't care what is passed through a channel. We only care when and if something is passed. In that situation, we can block and wait until something is sent on a channel using the following syntax.
```
<-ch
```
This will block until it pops a single item off the channel, then continue, discarding the item.

In cases like this, [Empty structs](https://dave.cheney.net/2014/03/25/the-empty-struct) are often used as a unary value so that the sender communicates that this is only a "signal" and not some data that is meant to be captured and used by the receiver.

Here is an example:
```
func downloadData() chan struct{} {
	downloadDoneCh := make(chan struct{})

	go func() {
		fmt.Println("Downloading data file...")
		// after the download is done, send a "signal" to the channel
		downloadDoneCh <- struct{}{}
	}()

	return downloadDoneCh
}

func processData(downloadDoneCh chan struct{}) {
	// any code here can run normally
	fmt.Println("Preparing to process data...")

	// block until `downloadData` sends the signal that it's done
	<-downloadDoneCh

	// any code here can assume that data download is complete
	fmt.Println("Data download ensured, starting data processing...")
}
```

### Buffered Channels
Channels can optionally be buffered.

#### Creating a channel with a Buffer
We can provide a buffer length as the second argument to `make()` to create a buffered channel like this:
```
ch := make(chan int, 100)
```

A buffer allows the channel to hold a fixed number of values before sending blocks. This mean sending on a buffered channel only blocks, when the buffer is full and receiving, blocks only when the buffer is empty.