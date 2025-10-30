package main

import (
	"fmt"
	"time"
)

type CustomWaitGroup struct {
	sem chan struct{}
}

// Add(n) увеличивает счётчик
func (w *CustomWaitGroup) Add(n int) {
	for i := 0; i < n; i++ {
		w.sem <- struct{}{}
	}
}

// Done() уменьшает счётчик
func (w *CustomWaitGroup) Done() {
	<-w.sem
}

// Wait() просит подождать
func (w *CustomWaitGroup) Wait() {
	for len(w.sem) > 0 {
		time.Sleep(10 * time.Millisecond)
	}
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		sem: make(chan struct{}, 100),
	}
}

func worker(id int, wg *CustomWaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	wg := NewCustomWaitGroup()

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, wg)
	}

	wg.Wait()
	fmt.Println("All workers finished!")
}
