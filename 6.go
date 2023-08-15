package main

import (
	"context"
	"fmt"
	"time"
)

func closeChan(in <-chan interface{}) {
	for range in {
	}
	fmt.Println("closeChan")
}

func cancelCtx(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("cancelCtx")
}

func chanWaitForRead(in <-chan interface{}) {
	<-in
	fmt.Println("chanWaitForRead")
}

func chanWaitForWrite(in chan<- struct{}) {
	in <- struct{}{}
	fmt.Println("chanWaitForWrite")
}

func main() {

	//1 close chan
	ch := make(chan interface{})
	go closeChan(ch)
	time.Sleep(time.Millisecond)
	close(ch)

	//2 cancel context
	ctx, cancel := context.WithCancel(context.Background())
	go cancelCtx(ctx)
	time.Sleep(time.Millisecond)
	cancel()
	time.Sleep(time.Millisecond)

	//3 read
	ch1 := make(chan interface{})
	go chanWaitForRead(ch1)
	time.Sleep(time.Millisecond)
	ch1 <- struct{}{}
	time.Sleep(time.Millisecond)

	//4 write to chan
	ch2 := make(chan struct{})
	go chanWaitForWrite(ch2)
	time.Sleep(time.Millisecond)
	<-ch2
	time.Sleep(time.Millisecond)

}
