package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 50; i < 60; i++ {
			ch2 <- i
		}
	}()
	go func() {
		defer close(ch3)
		for i := 65; i < 75; i++ {
			ch3 <- i
		}
	}()

	for v := range chanMerge(ch1, ch2, ch3) {
		fmt.Print(v)
	}
}

func chanMerge(channels ...<-chan int) <-chan int {
	out := make(chan int) // канал для объединенных данных
	var wg sync.WaitGroup

	wg.Add(len(channels)) // ждем завершения всех горутин
	for _, ch := range channels {
		go func(c <-chan int) {
			defer wg.Done()    // когда горутина завершилась завершаем ожидание
			for v := range c { // читаем из входного канала
				out <- v // отправляем в out канал
			}
		}(ch)
	}

	// когда все горутины завершатся закрываем out канал
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
