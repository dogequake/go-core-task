package main

import (
	"reflect"
	"testing"
)

// тест базовый кейс с пересечением
func TestIntersectSlices_Basic(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	expectedOk := true
	expectedSlice := []int{64, 3}

	gotOk, gotSlice := intersectSlices(a, b)

	if gotOk != expectedOk {
		t.Errorf("Ожидалось ok = %v, получено %v", expectedOk, gotOk)
	}

	if !reflect.DeepEqual(gotSlice, expectedSlice) {
		t.Errorf("Ожидалось срез = %v, получено %v", expectedSlice, gotSlice)
	}
}

// тест если второй слайс пустой то пересечения быть не может
func TestIntersectSlices_EmptySecond(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{}

	expectedOk := false
	expectedSlice := []int{}

	gotOk, gotSlice := intersectSlices(a, b)

	if gotOk != expectedOk {
		t.Errorf("Ожидалось ok = %v, получено %v", expectedOk, gotOk)
	}

	if !reflect.DeepEqual(gotSlice, expectedSlice) {
		t.Errorf("Ожидалось срез = %v, получено %v", expectedSlice, gotSlice)
	}
}

// тест если первый слайс пустой то пересечения быть не может
func TestIntersectSlices_EmptyFirst(t *testing.T) {
	a := []int{}
	b := []int{1, 2, 3}

	expectedOk := false
	expectedSlice := []int{}

	gotOk, gotSlice := intersectSlices(a, b)

	if gotOk != expectedOk {
		t.Errorf("Ожидалось ok = %v, получено %v", expectedOk, gotOk)
	}

	if !reflect.DeepEqual(gotSlice, expectedSlice) {
		t.Errorf("Ожидалось срез = %v, получено %v", expectedSlice, gotSlice)
	}
}

// тест если пересечений нет то результат пустой срез и false
func TestIntersectSlices_NoOverlap(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	expectedOk := false
	expectedSlice := []int{}

	gotOk, gotSlice := intersectSlices(a, b)

	if gotOk != expectedOk {
		t.Errorf("Ожидалось ok = %v, получено %v", expectedOk, gotOk)
	}

	if !reflect.DeepEqual(gotSlice, expectedSlice) {
		t.Errorf("Ожидалось срез = %v, получено %v", expectedSlice, gotSlice)
	}
}

// тест если все элементы пересекаются вернет все элементы
func TestIntersectSlices_AllOverlap(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{3, 2, 1}

	expectedOk := true
	expectedSlice := []int{1, 2, 3} // порядок как в a

	gotOk, gotSlice := intersectSlices(a, b)

	if gotOk != expectedOk {
		t.Errorf("Ожидалось ok = %v, получено %v", expectedOk, gotOk)
	}

	if !reflect.DeepEqual(gotSlice, expectedSlice) {
		t.Errorf("Ожидалось срез = %v, получено %v", expectedSlice, gotSlice)
	}
}
