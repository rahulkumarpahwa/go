package main

import (
	"fmt"
	"sync"
	"time"
)

func imageProcessingWorker(imageUrl string, wg *sync.WaitGroup, resultChannel chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 50) // trying to mimic the time taken to process the image.
	fmt.Println("image Processed : ", imageUrl)
	resultChannel <- imageUrl
}

func main() {
	var wg sync.WaitGroup
	resultChannel := make(chan string, 5)
	startTime := time.Now()

	wg.Add(5)
	go imageProcessingWorker("image1.png", &wg, resultChannel)
	go imageProcessingWorker("image2.png", &wg, resultChannel)
	go imageProcessingWorker("image3.png", &wg, resultChannel)
	go imageProcessingWorker("image4.png", &wg, resultChannel)
	go imageProcessingWorker("image5.png", &wg, resultChannel)

	wg.Wait()
	close(resultChannel) // channel always need to be closed, otherwise there will be deadlock.

	for val := range resultChannel {
		fmt.Println("Result of the Channel : ", val)
	}

	fmt.Printf("Estimated Time taken : %v", time.Since(startTime))

	// here we have sended the string in the channel but in actual we will send the struct or a map or even a slice.
}
