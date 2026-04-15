package main

import (
	"fmt"
	// "math/rand"
	"time"
	// "time"
)

// channel are like pipe which are used to pass data between the go routines when they are running concurrently.

// sending data to the processNum2 via channel
func processNum2(numChan *chan int) {

	// this will only take the one value and creates deadlock as many are sending but receiving only one.
	// fmt.Println("Processing Number ", <-*numChan)

	// solution of below problem in case 3:
	// this will able to receive multiple values which makes this work fine and available when multiple values are being send.
	for num := range *numChan {
		fmt.Println("Processing Number ", num)
		time.Sleep(time.Second) // just to mimic actual DB, file operation
	}
	// What range channel is already doing This line:
	// for num := range *numChan
	// already does blocking

}

// sum : will return the sum of the values passed to it.
// from Sum to main gooutine.
// func Sum(result chan int, num1 int, num2 int) {
// 	sum := num1 + num2
// 	result <- sum
// }

// goroutine synchronizer : (using channel as synchronizer)
// func task(done chan bool) {
// 	defer func() { done <- true }() // written in defer so as to get the called even if the code below fails
// 	fmt.Println("Process............")
// }

// // email queue system:
//
//	func emailSender(emailChan <-chan string, done chan<- bool) {
//		defer func() { done <- true }()
//		for email := range emailChan { // infinite loop, (keep waiting, until channel closed)
//			fmt.Println("Email has been sended to : ", email)
//			time.Sleep(time.Second) // mimic actual send
//		}
//	}
func main() {
	// ----------------------------------------------

	// CASE 7:
	// when want to listen at the multiple channels (say we have the two channel and we want to listen data once.)

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 5 // used by closure
	}()

	go func() {
		chan2 <- "Apple"
	}()

	for i := 0; i < 2; i++ {
		select { // helps to choose from both channel
		case chanVal1 := <-chan1:
			fmt.Println("received value from channel 1 : ", chanVal1)

		case chanVal2 := <-chan2:
			fmt.Println("received value from  channel 2 : ", chanVal2)
		}
	}

	// why loop around select ?
	// Iteration 1:
	// receives from one channel
	// Iteration 2:
	// receives from the other
	//  Both values are consumed

	// hence, here we are listining the both go channels at once.
	// similarly we can add more channels if we want.

	// --------------------------------------------------

	// CASE 6:
	// As of now, we are using the unbuffered channel, which has problem that it has sending and recieving blocking.
	// It means that when we have to send or recieve then we need to wait for goroutine to get completed to get data from channel.

	// So, we can use the buffered channels where we can send the data upto a limit a once without blocking and then can recieve unlike the unbuffered channel in which we have can send one data at time which we need to recieve back as well to send the data again through that.

	// we are here implementing queue system for email:

	// emailChan := make(chan string, 100) // 100 is the size of the buffer, upto which it is non blocking

	// done := make(chan bool)

	// emailChan <- "rahul@apple.com" // non-blocking
	// emailChan <- "kumar@apple.com" //non-blocking

	// fmt.Println(<-emailChan) // so, these will get known to main goroutine and execute without above being blocking, so no deadlock.
	// fmt.Println(<-emailChan)

	// ---------------------------------------------

	// CASE 6A: (this is simplest queue system we build)
	// so, now assume that we are recieving email from DB for which we need to send email. so, we will build a queue system for that.

	// emailChan := make(chan string, 100) // 100 is the size of the buffer, upto which it is non blocking

	// done := make(chan bool)

	// emailChan <- "rahul@apple.com" // non-blocking
	// emailChan <- "kumar@apple.com" //non-blocking
	// we will create the loop to send the multiple email at once as it is queue.

	//go emailSender(emailChan, done)
	// for i := 0; i <= 100; i++ {
	// 	emailChan <- fmt.Sprintf("%d@apple.com", i)
	// }
	// close(emailChan) //NOTE :  close because after the 100th email sended the in the method emailSender, there is inifnite loop, which is waiting for the emailChan to finished and then defer will call the done<-true so as to stop the main waiting here for the goroutine to get finished.
	// But, the main will never recieve the done as defer based line will never run, as stuck in the loop.
	// and '<-done' this is blocking when have no receiver will create the deadlock.

	// so, now here as all the calls in main are non-blocking then we need something to make it stop from
	// exiting before it executes the goroutines.
	// so, for that we will use another goroutine as the done synchronized as used in case 5.
	// <-done

	// -----------------------------------------

	// CASE 5:
	// previously we were using the waitgroup to make the main goroutine wait till all the goroutines are finished.
	// but same work can be done using the channels as well.

	// here we use the done boolean sended by the channel same as waitgroup.

	// when we have single goroutine then we can use the done chan as below, but when multiple gorotuine then we can use the waitgroup to do that.

	// done := make(chan bool)

	// go task(done)

	// status := <-done // blocking
	// fmt.Println("task Status : ", status)

	// ----------------------------------------------

	// // CASE 4:
	// result := make(chan int)

	// go Sum(result, 5, 6)

	// res := <-result // blocking
	// fmt.Println("Sum of the numbers from channel is :", res)
	// timeline of case4 :
	/*
		main:
		    start goroutine
		    wait on <-result   (BLOCKED)

		goroutine:
		    compute sum
		    send result

		→ handshake happens

		main:
		    resumes
		    prints result

	*/

	// -----------------------------------------------

	// CASE 3:

	// Normally we, don't send single numbers like numchan<-5, but as queue to send the large data.

	//numChan := make(chan int)

	//go processNum2(&numChan)

	// sending Number:
	// we will use here the infinite loop so as to create a queue to send multiple values.
	// we will think that 'This is also blocking and will make sure that our main will never stop.'
	// but in actual the blocking here is the sending itself (numChan<-rand.Intn(100)), because the recieving end has only capacity to recieve one value and no loop and when it recieve one value, it will end and then next (second iteration of the loop) iteration runs then sender statement will block again.
	// hence creating the deadlock as only sender is there but not receiver and we need both at the same time.

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	//--------------------------------------------------

	// CASE 2:

	// numChan := make(chan int)
	// go processNum2(&numChan)

	// // sending Number:
	// numChan <- 5
	// as this line above is blocking, it will block the main until the data is send to the recieving end, which is here the processNum2, but this does not guarantee the goroutine to get finished.
	// when data is recieved the blocking by this line is removed and then main (if it has nothing to do) will end
	// and if anything left which is to be excuted (like the goroutine) here. It will not get finished.

	// NOTE : Channel synchronization only guarantees data transfer, NOT goroutine completion

	// time.Sleep(time.Second)
	// so, as to stop the main to get completed before goroutine is started, because the sending via channel will be blocked until, the data is send. (it does not guarentee the goroutine completion)

	// ------------------------------------------------

	// CASE 1:
	// messageChan := make(chan string)

	// // sending data to the channel
	// // Note that the sending msg is blocking
	// // unless someone is recieving it on the other end of the pipe.
	// messageChan <- "ping"

	// // recieving data from channel
	// msg := <- messageChan

	// fmt.Println("data from channel", msg)

}

/*
IMPORTANT QUESTION :

func main()  {

	messageChan := make(chan string)

	// sending data to the channel
	// Note that the sending msg is blocking
	// unless someone is recieving it on the other end of the pipe.
	messageChan <- "ping"

	// recieving data from channel
	msg := <- messageChan

	fmt.Println("data from channel", msg)

}


if the channel is bloacking and it remains blocked unless someone is recieving then how here the blocked as we are receving as well?

Great catch — this is a classic confusion, and it means you’re really understanding how channels behave 👏

Let’s go step by step.

---

# 🧠 The key rule of channels

For an **unbuffered channel**:

> ✅ **Send blocks until a receiver is ready**
> ✅ **Receive blocks until a sender is ready**

---

# 🔥 What your code is doing

```go
messageChan := make(chan string)

messageChan <- "ping"   // ❌ BLOCKS HERE
msg := <-messageChan
```

---

# ⚠️ Important: This code will NOT work

Even though you wrote both send and receive…

👉 They are in the **same goroutine (main)**
👉 And execution is **sequential**

---

## 📉 What actually happens step-by-step

```text
1. Create channel
2. Try to send "ping"
3. No receiver exists yet ❌
4. Program BLOCKS here forever
```

👉 It never reaches:

```go
msg := <-messageChan
```

---

# 💥 Result

You’ll get:

```text
fatal error: all goroutines are asleep - deadlock!
```

---

# 🤔 Why this happens

Because Go does **NOT look ahead** and say:

> “Oh, a receiver is coming next line”

No ❌

It executes line-by-line:

* At send → must have receiver **right now**
* At receive → must have sender **right now**

---

# ✅ Correct way (using goroutine)

You need **two goroutines**:

```go
func main() {
	messageChan := make(chan string)

	go func() {
		messageChan <- "ping" // sender
	}()

	msg := <-messageChan // receiver

	fmt.Println("data from channel", msg)
}
```

---

## ⚙️ Now what happens

```text
main goroutine:
    waits to receive

new goroutine:
    sends "ping"

→ both meet → data passes → unblock
```

👉 This is called **synchronization point**

---

# 🧩 Think of it like a handshake

Unbuffered channel = handshake 🤝

* Sender: “I have data”
* Receiver: “I’m ready”
* Only when both meet → transfer happens

---

# ⚡ Buffered channel (extra insight)

If you use:

```go
messageChan := make(chan string, 1)
```

Then this works:

```go
messageChan <- "ping" // ✅ does NOT block
msg := <-messageChan
```

Why?

👉 Because buffer can hold 1 value

---

# 📊 Summary

| Case                        | Behavior                  |
| --------------------------- | ------------------------- |
| Unbuffered + same goroutine | ❌ Deadlock                |
| Unbuffered + goroutines     | ✅ Works                   |
| Buffered channel            | ✅ Works (within capacity) |

---

# 💡 Final Mental Model

> “Unbuffered channels require sender and receiver to meet at the same time.”

---

If you want, next I can show:

👉 How channels are used instead of WaitGroup
👉 OR
👉 How `select {}` works (super important for real systems)
*/

// -------------------------------------------

/* IMPORTANT TOPIC : SELECT
explain what is select here ?
explain what is select here ?
🧠 What is select in Go?

select lets you wait on multiple channel operations at the same time

👉 It’s like a switch for channels

⚙️ Your code
for i := 0; i < 2; i++ {
	select {
	case chanVal1 := <-chan1:
		fmt.Println("value of channel 1 : ", chanVal1)

	case chanVal2 := <-chan2:
		fmt.Println("value of channel 1 : ", chanVal2)
	}
}
🔥 What select actually does

At each iteration:

“Wait until ANY one of these channel operations is ready”

🔄 Step-by-step execution
1. Two goroutines start
chan1 <- 5
chan2 <- "Apple"

👉 Both are trying to send

2. First select iteration
select is waiting on:
    - receive from chan1
    - receive from chan2

👉 Whichever is ready first is picked

🎯 Important behavior

If BOTH are ready:

👉 Go picks randomly

Example outcomes
Run 1:
value of channel 1 : Apple
value of channel 1 : 5
Run 2:
value of channel 1 : 5
value of channel 1 : Apple

👉 Order is not guaranteed

🔁 Why loop runs 2 times?
for i := 0; i < 2; i++

👉 Because you expect 2 messages total

one from chan1
one from chan2
🧠 Mental Model

select is like:

“Listen to multiple doors 🚪”
Whichever door opens first → go there
⚡ Blocking behavior

select itself is blocking

👉 If none of the channels are ready:

select → waits (blocks)
🔥 Non-blocking select (extra concept)

If you add:

default:
	fmt.Println("no channel ready")

👉 Then:

It will NOT block
It will execute default immediately
🧩 Real-world usage

select is used for:

Handling multiple inputs
Timeouts
Cancellation (with context)
Event-driven systems
🚨 Small bug in your code
fmt.Println("value of channel 1 : ", chanVal2)

👉 Should be:

fmt.Println("value of channel 2 : ", chanVal2)
📊 Summary
Feature	Behavior
select	Waits on multiple channels
Blocking	Yes (unless default)
Multiple ready cases	Random selection
Loop	Needed to receive multiple values
💡 Final one-liner

“select waits on multiple channels and executes whichever becomes ready first.”
*/
