package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

/*

The Go program you've provided demonstrates the basic use of concurrency in Go through the use of goroutines.

Explanation of the Code:
The say function:

This function takes a string s as input and prints it 5 times with a 100ms delay between each print.
time.Sleep(100 * time.Millisecond) is used to introduce the delay.
The main function:

The main function launches a goroutine using the go keyword with the call go say("world").
This creates a concurrent execution of the say function with the argument "world".
After that, the main function continues execution and calls say("hello") directly (synchronously).
Key Concepts of Concurrency in Go:
Goroutines:

A goroutine is a lightweight thread managed by the Go runtime. It allows for concurrent execution of functions or methods.
The go keyword is used to start a goroutine. This enables concurrent execution of the say("world") function while the main function continues to execute the say("hello") function.
Each goroutine runs independently, and the Go runtime schedules them in a way that allows multiple goroutines to run concurrently.
Concurrency vs Parallelism:

Concurrency means that multiple tasks can start, run, and complete in overlapping time periods, but not necessarily simultaneously.
Parallelism means that tasks are literally running at the same time, which usually requires multiple CPU cores. Go can run goroutines concurrently, and if multiple CPU cores are available, Go can also execute them in parallel.
Synchronization:

In this program, there is no synchronization, meaning that the output from the goroutine and the main function can interleave. The Go runtime decides the order in which the goroutines are scheduled to run.
The say("hello") will most likely print first because it's running in the main thread, while the say("world") runs concurrently in the goroutine.

*/
