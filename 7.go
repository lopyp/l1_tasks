package main

import (
	"fmt"
	"sync"
)

func writer(id int, data *sync.Map, ch <-chan string) {
	for {
		key := <-ch
		data.Store(key, "value")
	}
}

func main() {
	// sync.Map map and mutex потокобезопасная map
	data := new(sync.Map)
	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("key-%d", i)
		}
		close(ch)
		wg.Done()
	}()

	go func() {
		for key := range ch {
			data.Store(key, "value")
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(data)
}
