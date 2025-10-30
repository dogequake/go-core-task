package main

import (
	"reflect"
	"testing"
)

// helper: собирает все значения из канала в срез
func collect(ch <-chan int) []int {
	var result []int
	for v := range ch {
		result = append(result, v)
	}
	return result
}

// тест сливание 3 канала
func TestChanMerge_Basic(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 2; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 10; i < 12; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(ch3)
		for i := 20; i < 22; i++ {
			ch3 <- i
		}
	}()

	got := collect(chanMerge(ch1, ch2, ch3))

	expected := []int{0, 1, 10, 11, 20, 21}

	if len(got) != len(expected) {
		t.Fatalf("Ожидалось %d элементов, получено %d", len(expected), len(got))
	}

	for _, val := range expected {
		found := false
		for _, g := range got {
			if g == val {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Элемент %d отсутствует в результате: %v", val, got)
		}
	}
}

// тест если один из каналов пустой
func TestChanMerge_OneEmpty(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
	}()

	go func() {
		defer close(ch2)
		ch2 <- 42
		ch2 <- 99
	}()

	got := collect(chanMerge(ch1, ch2))
	expected := []int{42, 99}

	if !reflect.DeepEqual(got, expected) && !reflect.DeepEqual(got, []int{99, 42}) {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}

// тест если все каналы пустые
func TestChanMerge_AllEmpty(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	close(ch1)
	close(ch2)
	close(ch3)

	got := collect(chanMerge(ch1, ch2, ch3))

	// nil и пустой срез — оба допустимы
	if len(got) != 0 {
		t.Errorf("Ожидалось пустой результат, получено %v", got)
	}
}

// тест с одним каналом
func TestChanMerge_SingleChannel(t *testing.T) {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()

	got := collect(chanMerge(ch))
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}
