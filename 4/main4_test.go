package main

import (
	"reflect"
	"testing"
)

// тест с задания
func TestSortSlice(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	expected := []string{"apple", "cherry", "43", "lead", "gno1"}
	result := sortSlice(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

// тест если второй слайс пустой то результат должен быть как первый слайс
func TestSortSlice_EmptySecond(t *testing.T) {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{}

	expected := []string{"a", "b", "c"}
	result := sortSlice(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

// тест если первый слайс пустой то результат должен быть пустой
func TestSortSlice_EmptyFirst(t *testing.T) {
	slice1 := []string{}
	slice2 := []string{"x", "y", "z"}

	expected := []string{}
	result := sortSlice(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

// тест если все элементы первого содержатся во втором то результат пустой
func TestSortSlice_AllExcluded(t *testing.T) {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"a", "b", "c", "d"}

	expected := []string{}
	result := sortSlice(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

// тест если элементы не пересекаются то результат должен быть как первый слайс
func TestSortSlice_NoOverlap(t *testing.T) {
	slice1 := []string{"x", "y", "z"}
	slice2 := []string{"a", "b", "c"}

	expected := []string{"x", "y", "z"}
	result := sortSlice(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}
