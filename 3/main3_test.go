package main

import (
	"testing"
)

// создаем StringIntMap для тестов
func newTestMap() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

// тест метода Add
func TestAdd(t *testing.T) {
	m := newTestMap()
	m.Add("Alice", 25)

	if _, ok := m.Get("Alice"); !ok {
		t.Errorf("ключ 'Alice' должен существовать после Add")
	}
}

// тест метода Exists
func TestExists(t *testing.T) {
	m := newTestMap()
	m.Add("Alice", 25)

	if !m.Exists("Alice") {
		t.Errorf("ключ 'Alice' должен существовать")
	}

	if m.Exists("Bob") {
		t.Errorf("ключ 'Bob' не должен существовать")
	}
}

// тест метода Get
func TestGet(t *testing.T) {
	m := newTestMap()
	m.Add("Alice", 25)

	val, ok := m.Get("Alice")
	if !ok || val != 25 {
		t.Errorf("ожидалось (25, true), получили (%d, %v)", val, ok)
	}

	_, ok = m.Get("Bob")
	if ok {
		t.Errorf("ключ 'Bob' не должен существовать")
	}
}

// тест метода Remove
func TestRemove(t *testing.T) {
	m := newTestMap()
	m.Add("Alice", 25)
	m.Remove("Alice")

	if m.Exists("Alice") {
		t.Errorf("ключ 'Alice' должен быть удален")
	}
}

// тест метода Copy
func TestCopy(t *testing.T) {
	m := newTestMap()
	m.Add("Alice", 25)
	m.Add("Bob", 30)

	copyMap := m.Copy()

	for k, v := range m.data {
		if copyMap[k] != v {
			t.Errorf("копия некорректна для ключа %s", k)
		}
	}

	copyMap["Alice"] = 100
	if m.data["Alice"] == 100 {
		t.Errorf("изменение копии повлияло на оригинал")
	}
}
