package main

import (
	"testing"
	"time"
)

// тест базовый функционал CustomWaitGroup
func TestCustomWaitGroup(t *testing.T) {
	wg := NewCustomWaitGroup()

	start := time.Now()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)

	if elapsed < 100*time.Millisecond {
		t.Errorf("Wait завершился слишком рано: %v", elapsed)
	}
}

// тест если все горутины сразу вызывают доне
func TestCustomWaitGroup_AllDone(t *testing.T) {
	wg := NewCustomWaitGroup()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(50 * time.Millisecond)
			wg.Done()
		}()
	}

	wg.Wait()

	if len(wg.sem) != 0 {
		t.Errorf("После Wait канал должен быть пуст, а в нём %d элементов", len(wg.sem))
	}
}
