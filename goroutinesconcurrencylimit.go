package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	outputLimit := make(chan int, 4)
	var waitGroup sync.WaitGroup
	for index := 0; index < 20; index++ {
		waitGroup.Add(1)
		outputLimit <- index
		go process(index, outputLimit, &waitGroup)
	}
	waitGroup.Wait()
	fmt.Println("Finished")
}

func process(index int, outputLimit chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Processing: ", index)
	time.Sleep(time.Duration(time.Second * 3))
	<-outputLimit
}
