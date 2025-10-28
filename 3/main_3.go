package main

import (
	"fmt"
)

type StringIntMap struct {
	data map[string]int
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	//вообще сюда по хорошему нужно добавить проверку на существование ключа
	//но в задании этого не требуется
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int, len(m.data))
	for k, v := range m.data {
		newMap[k] = v
	}
	return newMap
}

func (m *StringIntMap) Exists(key string) bool {
	if _, ok := m.data[key]; ok {
		return true
	}
	return false
}

func (m *StringIntMap) Get(key string) (int, bool) {
	//здесь тоже можно сделать проверку на существование ключа
	//но в задании этого не требуется
	value, ok := m.data[key]
	return value, ok
}

func main() {
	fmt.Println("Это 3 задание про StringIntMap)")
}
