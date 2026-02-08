package main

import (
	"fmt"
	"sync"
)

// wait group : mechanism, in the previous part of the goroutines we see that until all of our go routines gets executed we need to pass the sleep in the main function.
// it is okay for the example but not good for the real world as we don't know how much time it would take to get completed.
// to handle that we use the waitgroup.

// we want to synchronous the way that when all the go routines get completed then only the main method should completed. for that we use the wait group.

func task(id int, w *sync.WaitGroup) { // we have to pass the waitgroup as the pointer here from the sync package.

	// after the task has been done we have to remove the waitgroup which is done with 'Done()' method with the use of defer keyword. this 'defer' makes sure that the line runs at the end with the defer keyword.
	defer w.Done()
	// this will minus the 1 we have added while making the call in the main.
	fmt.Println("Doing task with id : ", id)
}

func main() {
	var wg sync.WaitGroup
	// A WaitGroup is a counting semaphore typically used to wait for a group of goroutines or tasks to finish.

	for i := 0; i <= 10; i++ {
		wg.Add(1) // this waitgroup wg acts as a counter so we will add the one in it.
		// we added one because we start one goroutines. as the next go routines starts one will be added. so, each time one go rotines add one.
		// the above defer keyword Done() method will remove that one.
		go task(i, &wg)
		// we pass here the pointer of the waitgroup so as to get call the Done() with the defer.
	}

	wg.Wait() //here we have to do.
	// so, program will wait until our waitgroup not gets to zero by-default.

	// so in waitgroup, there are three things : 1. Done, 2. Wait 3. Add

	//---- when used with anonmous method then we don't need to pass the wg in the that as we call directly call as:
	var wg2 sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func() {
			fmt.Println("doing task by anonmous fxn: ", i)
			wg2.Done()
		}()

	}
	wg2.Wait()

}
