package main

import (
	"fmt"
)

func convertUint8ToFloat64(in <-chan uint8) <-chan float64 {
	out := make(chan float64)
	go func() {
		for v := range in {
			out <- float64(v)
		}
		close(out)
	}()
	return out
}

func cubeNumbers(in <-chan float64) <-chan float64 {
	out := make(chan float64)
	go func() {
		for v := range in {
			out <- v * v * v
		}
		close(out)
	}()
	return out
}

func main() {
	ch1 := make(chan uint8)

	// Горутинa, которая записывает числа в первый канал
	go func() {
		defer close(ch1)
		for i := uint8(1); i <= 10; i++ {
			ch1 <- i
		}
	}()

	// Конвейер: ch1 → convertUint8ToFloat64 → cubeNumbers
	ch2 := cubeNumbers(convertUint8ToFloat64(ch1)) // ✅ ch2 локальная, тип <-chan float64

	// Читаем результаты из конечного канала и выводим
	for result := range ch2 {
		fmt.Println(result)
	}
}
