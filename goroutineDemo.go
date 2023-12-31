package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("- - v231230 test golang demo program - - - -\n")

	// use make() to Create a channel for communication between Goroutines
	dataChannel := make(chan int)

    // Use a WaitGroup to wait for both Goroutines to finish
    // I believe this is just an integer count of the number of channels open? And on close, it may decrement?
    var wg sync.WaitGroup

    // Goroutine 1: Sending data to the channel
    // add to list of open waitgroup items?
    // seems to be so, cause if comment out get panic: sync: negative WaitGroup counter
    // but if try to view variable get something odd like &{{} {{} {} 4294967297} 0}
    wg.Add(1)
    // use go to run as a goroutine. Pass in channel and ref of wg variable into the function
    go sendData(dataChannel, &wg)

    // Goroutine 2: Receiving data from the channel
    wg.Add(1)
    go receiveData(dataChannel, &wg)

	fmt.Println("proof that the main program goes past the gorotuines aka background jobs and is not blocked by them\n")

    // Wait for both Goroutines to finish
    wg.Wait()

	fmt.Println("- - End of program, this print statement waited for the waitgroup to finish - - - -")
}

func sendData(c chan int, wg *sync.WaitGroup) {
// Ensure that the WaitGroup is decremented when the Goroutine finishes
defer wg.Done()

//for 1-5, so 5x, just send number to channel, wait 0.5 sec
for i := 1; i <= 5; i++ {
// Send data to the channel
    	fmt.Println("Sending the data:", i) //, "wg=", wg)
c <- i
time.Sleep(time.Millisecond * 500) // Simulate some work
}

    // Close the channel when done sending data
    //generally, it is the sender's job to close the channel when done
    close(c)

}

func receiveData(c chan int, wg *sync.WaitGroup) {
// Ensure that the WaitGroup is decremented when the Goroutine finishes
defer wg.Done()

    //special code, for every item/input received from this channel c, every time data is received, run the for loop. Should close once the input channel is closed
    // Receive data from the channel until it's closed
    for data := range c {
    	fmt.Println("Received:", data)
    	fmt.Println("Received:", data, " and 3x then +1 is: ", data * 3 + 1, "\n")
    }

}