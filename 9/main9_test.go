package main

import (
	"reflect"
	"testing"
)

func collectFloat64(ch <-chan float64) []float64 {
	var result []float64
	for v := range ch {
		result = append(result, v)
	}
	return result
}

// базовый тест для convertUint8ToFloat64

func TestConvertUint8ToFloat64_Basic(t *testing.T) {
	in := make(chan uint8)
	go func() {
		defer close(in)
		in <- 1
		in <- 2
		in <- 3
	}()

	got := collectFloat64(convertUint8ToFloat64(in))
	expected := []float64{1, 2, 3}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}

// тест пустой convertUint8ToFloat64
func TestConvertUint8ToFloat64_Empty(t *testing.T) {
	in := make(chan uint8)
	close(in)

	got := collectFloat64(convertUint8ToFloat64(in))

	if len(got) != 0 {
		t.Errorf("Ожидалось пустой результат, получено %v", got)
	}
}

// базовый тест для cubeNumbers

func TestCubeNumbers_Basic(t *testing.T) {
	in := make(chan float64)
	go func() {
		defer close(in)
		in <- 1
		in <- 2
		in <- 3
	}()

	got := collectFloat64(cubeNumbers(in))
	expected := []float64{1, 8, 27}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}

// тест пустой cubeNumbers
func TestCubeNumbers_Empty(t *testing.T) {
	in := make(chan float64)
	close(in)

	got := collectFloat64(cubeNumbers(in))

	if len(got) != 0 {
		t.Errorf("Ожидалось пустой результат, получено %v", got)
	}
}

// тест конвейера на преобразование и возведение в куб

func TestPipeline_Full(t *testing.T) {
	in := make(chan uint8)
	go func() {
		defer close(in)
		for i := uint8(1); i <= 5; i++ {
			in <- i
		}
	}()

	out := cubeNumbers(convertUint8ToFloat64(in))

	got := collectFloat64(out)
	expected := []float64{1, 8, 27, 64, 125}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}
