# Mutexes in GO
Mutexes allow us to lock access to data. This ensures that we can control which goroutines can access certain data at which time.

Go's standard library provide a  built-in implementation of a mutex with the [sync.Mutex](https://pkg.go.dev/sync#Mutex) type and its two methods:
- [Lock()](https://golang.org/pkg/sync/#Mutex.Lock)
- [Unlock()](https://golang.org/pkg/sync/#Mutex.Unlock)

We can protect a block of code by surrounding it with a call to `Lock` and `Unlock` as shown on the `protected()` function below.

It's a good practice to structure the protected code within a function so that `defer` can be used to ensure that we never forget to unlock the mutex
```
func protected() {
    mu.Lock()
    defer mu.Unlock()
    // The rest of the function is protected any other calls to `mu.Lock()` will block
}
```

Mutexes are powerful. Like most powerful things, they can also cause many bugs if used carelessly.

### Maps are not [thread-safe](https://en.wikipedia.org/wiki/Thread_safety)
Maps are `not` thread-safe for concurrent use. If we have multiple goroutines accessing the same map, and at least one of them is writing to the map, we must [lock](https://en.wikipedia.org/wiki/Readers%E2%80%93writer_lock) the maps with a mutex.