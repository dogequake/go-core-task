package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomGenerator(ch chan int, count int) {
	for i := 0; i < count; i++ {
		num := rand.Intn(100)
		ch <- num
	}
	close(ch)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int) // небуферизированный канал

	go randomGenerator(ch, 10)

	for num := range ch {
		fmt.Println(num)
	}
}
