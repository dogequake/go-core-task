package main

import (
	"testing"
)

// тест канал закрывается корректно после генерации
func TestRandomGenerator_ChannelClosed(t *testing.T) {
	ch := make(chan int)
	count := 3

	go randomGenerator(ch, count)

	for i := 0; i < count; i++ {
		_, ok := <-ch
		if !ok {
			t.Errorf("Канал закрылся раньше времени")
		}
	}

	_, ok := <-ch
	if ok {
		t.Errorf("Канал не закрылся после генерации всех чисел")
	}
}
