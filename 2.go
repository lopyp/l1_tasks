package main

import (
	"fmt"
	"sync"
)

func main() {
	data := []int{2, 4, 6, 8, 10}
	var wait sync.WaitGroup
	resultCh := make(chan int)
	for _, i := range data {
		wait.Add(1)
		go square(i, &wait, resultCh)
	}
	go func() {
		wait.Wait()
		close(resultCh)
	}()
	for res := range resultCh {
		fmt.Printf("%d\n", res)
	}
}

func square(num int, wait *sync.WaitGroup, resultCh chan<- int) {
	defer wait.Done()
	res := num * num
	resultCh <- res
}
