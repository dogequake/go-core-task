package main

import (
	"testing"
)

// тест sliceExample
func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}
	result := sliceExample(input)
	if len(result) != len(expected) {
		t.Errorf("ожидалось %v, получили %v", expected, result)
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("на позиции %d ожидалось %d, получили %d", i, expected[i], result[i])
		}
	}
}

// тест addElements
func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	element := 5
	expected := []int{1, 2, 3, 5}
	result := addElements(input, element)
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("на позиции %d ожидалось %d, получили %d", i, expected[i], result[i])
		}
	}
}

// тест copySlice
func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3}
	result := copySlice(input)
	if len(result) != len(input) {
		t.Errorf("длина копии %d, ожидалось %d", len(result), len(input))
	}
	for i := range input {
		if result[i] != input[i] {
			t.Errorf("на позиции %d ожидалось %d, получили %d", i, input[i], result[i])
		}
	}
	result[0] = 100
	if input[0] == result[0] {
		t.Errorf("изменение копии повлияло на оригинал")
	}
}

// тест removeElement
func TestRemoveElement(t *testing.T) {
	input := []int{10, 20, 30, 40, 50}
	index := 2
	expected := []int{10, 20, 40, 50}
	result := removeElement(input, index)
	if len(result) != len(expected) {
		t.Errorf("длина %d, ожидалось %d", len(result), len(expected))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("на позиции %d ожидалось %d, получили %d", i, expected[i], result[i])
		}
	}
}
