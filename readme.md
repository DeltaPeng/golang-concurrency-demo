https://gobyexample.com/
https://gophercises.com/

to install Go, first get an IDE, like VS code, then go to official golan page go.dev
get installer, run it. Exit command line window, re-enter, should be able to type in "go" command and now see help options

to create a project, just
mkdir your-app-name
cd your-app-name
code .

in golang, functions that start with capital letter are public, if they start with lower they are private to the package

if you want an .exe, create a main package with a main function. This is like c++, it starts in main, and ends when it gets to the end of main. It's the main program/goroutine/thread, it's procedral. If you are creating a library kind of package, then DON'T have a main package with it.

Goroutines are like running a job in the background. A function can be run as a goroutine by running "go" command right before it. A function may take a channel as an input parameter, i.e chan int is a channel whose data type of data it passes around is an int. A function that has a channel is not necessarily a goroutine nor does it always have to be run as such. A channel is a means to pass around data safely/securely between goroutines, like a subscribe/notify kind of set up. There's a special loop code that will run each time it receives data from the channel.

If a variable has \* at the front of it, it means it's the reference to that variable (the pointer), or if the value is a pointer, it can be used to dereference aka get the actual value at the variable. A & on a variable gives the pointer of the variable instead of it's value, so, on a function def you may have \*, but then a & when passing an input param into it, i.e.

//in main code
// use go to run as a goroutine. Pass in channel and ref of wg variable into the function
go sendData(dataChannel, &wg)

//in func def, expects a pointer
func sendData(c chan int, wg \*sync.WaitGroup) {
...

if you have defer, it means runs this function on close, to help with cleanup regardless of how a function exits (whether it closes gracefully or not)

Golang is useful because it is better in more easily handles multithreading (was designed for it), to take advantage of modern hardware. Can handle 1000's? ~unlimtied? threads safely and efficiently. Programs like Docker, Hashicorp Vault are written in GoLang. Made by google in 2009. Can infer types. Faster than interpretative languages, cause it compiles to a binary code. Can run on windows, mac, linux. Compiles quickly. Primarily used for backend or server type coding, microservices, apis. Not as much so for ui's or gui's.

sample structure
myproject/
├── main.go
├── pkg/
│ └── mypackage/
│ └── mypackage.go
└── src/
└── myproject/
├── mysubpackage/
│ └── mysubpackage.go
└── mylibrary/
└── mylibrary.go

once a project is done, if it has a main package, can run "go build main.go" to compile it to binary code. For windows it will package code with dependencies to create an .exe file, of the same name as the .go file with the main package

----gocode can be passed into a docker container-------
, and the command set to run it (similar to a .sh program). This method below is passing the source code in, and building it on the docker container, not safe way to do it for production, as someone could steal the code if they have access to the docker container

FROM golang:latest

WORKDIR /app
COPY . .

RUN go build -o myapi

CMD ["./myapi"]

instead could build the app, then only copy that in

# Use a minimal base image for the runtime

FROM alpine:latest

# Set the working directory inside the container

WORKDIR /app

# Copy only the built binary from the local machine to the container

COPY myapi .

# Expose the port that your Go application listens on

EXPOSE 8080

# Specify the command to run when the container starts

CMD ["./myapi"]

--Golang code example --------
package main

import (
"fmt"
"sync"
"time"
)

func main() {
// use make() to Create a channel for communication between Goroutines
dataChannel := make(chan int)

    // Use a WaitGroup to wait for both Goroutines to finish
    // I believe this is just an integer count of the number of channels open? And on close, it may decrement?
    var wg sync.WaitGroup

    // Goroutine 1: Sending data to the channel
    // add to list of open waitgroup items?
    wg.Add(1)
    // use go to run as a goroutine. Pass in channel and ref of wg variable into the function
    go sendData(dataChannel, &wg)

    // Goroutine 2: Receiving data from the channel
    wg.Add(1)
    go receiveData(dataChannel, &wg)

    // Wait for both Goroutines to finish
    wg.Wait()

}

func sendData(c chan int, wg \*sync.WaitGroup) {
// Ensure that the WaitGroup is decremented when the Goroutine finishes
defer wg.Done()

//for 1-5, so 5x, just send number to channel, wait 0.5 sec
for i := 1; i <= 5; i++ {
// Send data to the channel
c <- i
time.Sleep(time.Millisecond \* 500) // Simulate some work
}

    // Close the channel when done sending data
    //generally, it is the sender's job to close the channel when done
    close(c)

}

func receiveData(c chan int, wg \*sync.WaitGroup) {
// Ensure that the WaitGroup is decremented when the Goroutine finishes
defer wg.Done()

    //special code, for every item/input received from this channel c, every time data is received, run the for loop. Should close once the input channel is closed
    // Receive data from the channel until it's closed
    for data := range c {
    	fmt.Println("Received:", data)
    }

}
