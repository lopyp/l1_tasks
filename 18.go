package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := &Counter{}

	var wg sync.WaitGroup

	// Запускаем 10 горутин, которые будут инкрементировать счетчик
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000000; j++ {
				counter.Increment()
			}
			wg.Done()
		}()
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println(counter.Value())
}
