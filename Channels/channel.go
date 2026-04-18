package main

// In Go, a channel is like a pipe through which goroutines can send and receive data safely.
// 👉 Think of it like:
// Two people passing a box through a tube — one puts data in, the other takes it out.In Go, a channel is like a pipe through which goroutines can send and receive data safely.



// Why do we need Channels?
// Before channels, sharing data between threads (or goroutines) caused problems like:
// Data corruption
// Race conditions
// Complex locks (mutex headache 😵)
// Go’s philosophy:
// “Don’t communicate by sharing memory; share memory by communicating.”
// Channels help you:
// Avoid manual locking (like sync.Mutex)
// Safely pass data between goroutines
// Synchronize execution

import "fmt"

func main(){


	// Channels are used for 1) holds data  2)thread safe 3) listen to data

	// var c = make(chan int) // this channel can hold int value
    //  c<-1   // means we add value 1 to channel
	//  var ch= <- c  // retriving value from channels
	//  fmt.Println(ch)

	// this above code gives deadlock becuase there is no one to recive the value 1 it keeps waiting and give dealock error

	

    c := make(chan int)

    go func() {
        c <- 5   // send
    }()

    num := <-c  // receive

    fmt.Println(num)




}