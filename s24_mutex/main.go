package main

import (
	"fmt"
	"sync"
)

// mutex : mutual exclusion
// when we use the multi threading and to save outself from the race condition then use the mutex.

// race condition: when multiple processes try to modify / write to same resource then operation will not be atomic and it will be a race condition.
// eg, when one process done some changes to a resource then other process overwrite it then it is race condition.

// eg, we are making a social media app and it has posts which has views and has struct as below:

type post struct {
	views int
	// we will normally create the mutex here in the struct itself. other we can create it glovbally as well in the main.
	mu sync.Mutex
}

// now, we have to increament the views. so we will create the method attach to it.
func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()

	// now, where the modification is done, we will lock with mutex and after that we will unlock.
	p.mu.Lock()
	p.views += 1
	// p.mu.Unlock()
	// now, when one go routine is working on it, then other go routines will wait till it gets unlock and then will work on it.
	// we can also put the unlock in the defer method as well (best practice). as the modification logic may have the error and in that case we will not able to reach the unlock statement.
}

func main() {
	var wg sync.WaitGroup

	myPost := post{views: 0}
	// myPost.inc()
	// myPost.inc()

	// but, in real world apps, all the tasks done are not done one by one but multiple request comes to update the views concurrently.
	// so, in such cases we will use the goroutines.

	// so, in case of asynchronous we will use the go and let's it update 100 times and to get the goroutine gets completed we will use the waitgroup so as make the main wait till completion.

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()

	// now we will have the different value of the views, some time 100, sometime 99, and sometime other value.
	fmt.Println("My Post Views : ", myPost.views)
	// this arbitary value in the views is due to concurrency we used with goroutine.
	// because there is race condition which is caused by the light weight threads which are trying to modify the views at the same time.
	// this caused to override the changes done by one process, by another process.
	// which in result the not the actual value but arbitary value.
	//so , we will use the mutex to lock the modification, so that i can be done only by one process at a time.

}
