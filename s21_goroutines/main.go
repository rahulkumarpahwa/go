package main

import (
	"fmt"
	"time"
)

// go routines are the light weight threads when you have to do multithreading and run the things concurrently can be done by the go routines.

// eg;. we have the task method which takes id and does the task.

func task(id int) {
	fmt.Println("doing task", id)
}

func main() {
	for i := 0; i < 10; i++ {
		task(i)
	}
	// here the method task is called for each id one by one. first for the id =1, then id =2 and so on. so the method is blocking. means call can be done to one at a time.

	// so to make unblocking we have to use the goroutines which will be done by calling the method in the loop concurrently.
	// to do that just write the 'go' in front of the method.

	for i := 0; i < 10; i++ {
		go task(i)
	}
	// so, now each task will be performed in the light weight thread which is called go routines, concurrently. so there will called parallely based upon the no. of cores in the system.
	// also NO log will be there as the main function is also running in the go routines, which is called also.
	// here in the case, when the 10 goroutines are created and called then along our main function will be called which after which no time is left for the goroutines to get finished as the main func finishes so that's why the goroutines are left un printed and pending gorotines are exited inbetween.

	// so, to get the goroutines finsihes as the main finishes we will make it sleep for few seconds.
	fmt.Println("hello World!")
	time.Sleep(time.Second * 1)
	// this will make the go routines get completed.
	// also the order of the goroutines printed is not same as one in normal task printed one by one, which means they are fast and non blocking.

	// we can also do the anonmous function as well:

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("doing task by anonmous method", i) // applying the closure here for the i and it is anonmous method.
		}()
	}

	time.Sleep(time.Second * 1)
	// normally we don't put the sleep like this in the main func, we have the mechanism proper for that as well.

}
