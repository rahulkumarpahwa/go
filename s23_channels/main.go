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

// goroutine synchronizer
// we have taken the bool channel here in the function
func task(done chan bool) {
	defer func() { done <- true }() // this will call after the function call even when there is error in the method task.defer is mpstly used with function call.
	// ai definition of defer : The defer keyword in Go is used to schedule a function call to run just before the surrounding function returns, making it a powerful tool for ensuring cleanup actions are executed reliably.
	fmt.Println("processing.....")
}

// we will use this buffered channel.
func sendEmail(emailChannel chan string) {

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

	//18:00

	done := make(chan bool)
	go task(done)

	// now to stop until our goroutines are completed other than main, so we will need something blocking and we know that the sending and reciving part of the channel are blocking.
	<-done //block, this recoieving value is block, we will not need to use the bool value returned here as the channel 'done' when get the value true in the task method which is goroutine and this will make the <-done part gets unblocked. the bool true passed to the done when the gorotuine are completed in the task as the defer will make the call at the end.

	// so, we can do that completion of the goroutine with the channel as well other than goroutine.
	// mostly, we will use the channel when we have a single goroutine then we use this channel otherwise for the multiple task use the waitgroup is more useful as it has the counter.

	//23:51
	// until now, we have seen the unbuffered channel. the problem is with the sending and recieving with it, is blocking. until, sending and recieving is being processed, we have to wait. but, if we have to build something queue system where everything is send immediately. then we can use the buffered channel.
	// in unbuffered channel, we can send the data from the one goroutine to another but the buffered one is blocking. ie. it is send one by one
	// But in buffered, we can send the limited amount of data without blocking.

	// we will see a queue implmentation, for a email using the buffered channel.

	emailChannel := make(chan string, 100) // in IRL, we will send the structure as we have to send the metadata as well.
	// here we are passing the size for the buffered channel as well as the second parameter.

	emailChannel <- "rahul@apple.com" // sending data to channel, non blocking.
	emailChannel <- "rk@apple.com"    // sending data to channel, non blocking.

	fmt.Println("email channel : ", <-emailChannel) // this will get printed and not get the error as we have used the buffered channel and till the memory is finished, this will be not blocked.
	fmt.Println("email channel : ", <-emailChannel) // this will get printed and not get the error as we have used the buffered channel and till the memory is finished, this will be not blocked.

	// now we will apply the queue to send the email to the users for which we are gettting from the DB and we have to send them in batch with emailSender method we will make ourself as:
	// here we will use the two channels one to send and other to track the gorotuines are completed then mains should be completed.
	// 27:10

}
