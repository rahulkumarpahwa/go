package main

import (
	"fmt"
	"sync"
	"time"
)

func imageProcessingWorker(imageUrl string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 50) // trying to mimic the time taken to process the image.
	fmt.Println("image Processed : ", imageUrl)
}

func main() {

	startTime := time.Now()
	var wg sync.WaitGroup

	wg.Add(5)
	go imageProcessingWorker("image1.png", &wg)
	go imageProcessingWorker("image2.png", &wg)
	go imageProcessingWorker("image3.png", &wg)
	go imageProcessingWorker("image4.png", &wg)
	go imageProcessingWorker("image5.png", &wg)

	wg.Wait()
	fmt.Printf("Estimated Time taken : %v", time.Since(startTime))
}
