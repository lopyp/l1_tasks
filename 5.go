package main

import (
	"fmt"
	"time"
)

func main() {
	var N int
	_, err := fmt.Scanf("%d", &N)
	if err != nil {
		fmt.Print("error time")
	}

	//канал для передачи данных
	dataCh := make(chan int)
	//горутинa для записи данных в канал
	go func() {
		i := 0
		for {
			dataCh <- i
			i++
		}
	}()

	//горутинa для чтения данных из канала
	go func() {
		for data := range dataCh {
			fmt.Printf("Received data: %d\n", data)
		}
	}()

	time.Sleep(time.Duration(N) * time.Second)
	close(dataCh)
}
