package main

import (
	"fmt"
	// "math/rand"
	"time"
)

// channels : these are like the pipes in which one side we enter data and other side we get it.
// we use this with go-routines when we have to send the data from one goroutine to another.

// for sending : from main to goroutine.
func processNum(numChan chan int) {
	// fmt.Println("processing number : ", <-numChan)
	// now values are coming in the multiple so we need to have loop to print them as well. so we will do that using the range.
	for num := range numChan {
		fmt.Println("processing number : ", num)
		time.Sleep(time.Second * 1)
		// to get the value printed we need sleep.
	}
}

// for sending to main goroutine:
func sum(result chan int, num1 int, num2 int) {
	result_sum := num1 + num2
	result <- result_sum // here the result is the channel and result is just the sum of two numbers.

}

func main() {
	// messageChan := make(chan string)
	// simply create the variable and use the make method to pass to make the 'chan' keyword and then tell the type which is of the channel we made.

	// messageChan <- "ping" // we are sending the data into the channel. this arrow toward the channel shows we are sending the data towards the channel.

	// now to recieve we have put the arrow outside the variable of channel and take the value in some another variable.
	// msg := <-messageChan

	// fmt.Println("message through the channel : ", msg)
	// here we get the error of DEADLOCK :
	/* fatal error: all goroutines are asleep - deadlock!
	goroutine 1 [chan send]:
	main.main()
	C:/Programming/projects/golang/s23_channels/main.go:12 +0x36
	exit status 2
	*/

	// now we get the deadlock because the CHANNEL are blocking unless someone is there to recieve the code otherside.
	// so, to use them properly:

	// numChan := make(chan int)
	// we will also create the method which we call and pass the channel in there.

	// go processNum(numChan) // here we are creating the channel between the main goroutine and the processNum go routine.

	// numChan <- 5 // passing the number 5 in the channel numChan.

	// we need make the main sleep as the processNum is non-blocking and this will run the passing data to the numChan immediately which will make us print nothing so,

	// time.Sleep(time.Second * 1)

	//so, we see that how to send the data from the one goroutine to another using the channel.

	//10:31

	// but normally we don't send data like this for the one value. 
	// we send the data like the queue using the loop.

	// for { // for loop inifinte ie. without condition
	// 	numChan <- rand.Intn(100) // this generates the random number between the given range.
	// }

	// time.Sleep(time.Second * 1) //no need tpo put this as our program does not stop under inifite loop.

	// we will use the loop in fxn as well. (See above)

	// this can be powerfull, as this is used to apply queue system and synchronise different threads.

	//13:53
	// now, we can send back data from the goroutine (other than main) to main goroutine via channel as:

	result := make(chan int) // new channel of time int

	go sum(result, 3, 4)

	result_recieved := <-result // this reading is blocking. so we don;t need sleep here.
	fmt.Println("getting result from the goroutine to the main goroutine : ", result_recieved)

	//18:32

}
