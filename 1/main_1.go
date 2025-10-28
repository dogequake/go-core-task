package main

import (
	"fmt"
)

var numDecimal int = 42
var numOctal int = 052
var numHexadecimal int = 0x2A
var pi float64 = 3.14
var name string = "Golang"
var isActive bool = true
var complexNum complex64 = 1 + 2i

func main() {
	// определение типа каждой переменной и вывод его на экран
	definition_type()

	// преобразование всех переменных в строку
	str := convert_to_string()
	fmt.Printf("Преобразованные переменные в строковый тип: %q\n", str)

	// преобразование строки в срез рун
	runes := convert_to_rune_slice(str)
	fmt.Printf("Срез рун: %q\n", runes)

	// добавление соли и вывод итоговой строки
	fmt.Printf("Итоговая строка с солью: %q\n", rune_with_salt(runes))
}

func definition_type() {
	fmt.Printf("Тип numDecimal: %T\n", numDecimal)
	fmt.Printf("Тип numOctal: %T\n", numOctal)
	fmt.Printf("Тип numHexadecimal: %T\n", numHexadecimal)
	fmt.Printf("Тип pi: %T\n", pi)
	fmt.Printf("Тип name: %T\n", name)
	fmt.Printf("Тип isActive: %T\n", isActive)
	fmt.Printf("Тип complexNum: %T\n", complexNum)
}

// Альтернативный вариант с использованием пакета reflect
// func definition_type() {
// 	fmt.Printf(reflect.TypeOf(numDecimal).String())
// 	fmt.Printf(reflect.TypeOf(numOctal).String())
// 	fmt.Printf(reflect.TypeOf(numHexadecimal).String())
// 	fmt.Printf(reflect.TypeOf(pi).String())
// 	fmt.Printf(reflect.TypeOf(name).String())
// 	fmt.Printf(reflect.TypeOf(isActive).String())
// 	fmt.Printf(reflect.TypeOf(complexNum).String())
// }

func convert_to_string() string {
	str := fmt.Sprintf("%d%d%d%.2f%s%t%v",
		numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	return str
}

// Альтернативный вариант с использованием пакета strconv
// func convert_to_string() {
// 	str := strconv.Itoa(numDecimal) +
// 		strconv.Itoa(numOctal) +
// 		strconv.Itoa(numHexadecimal) +
// 		strconv.FormatFloat(pi, 'f', 2, 64) +
// 		name +
// 		strconv.FormatBool(isActive) +
// 		strconv.FormatComplex(complex128(complexNum), 'f', 2, 64)
// 	fmt.Printf("Преобразованные переменные в строковый тип: %q\n", str)
// }

func convert_to_rune_slice(s string) []rune {
	runeSlice := []rune(s)
	return runeSlice
}

func rune_with_salt(runes []rune) string {
	// соль
	salt := []rune("go-2024")

	// вставка соли в середину среза рун
	mid := len(runes) / 2
	runesWithSalt := append(runes[:mid], append(salt, runes[mid:]...)...)

	// превращаем обратно в строку
	finalString := string(runesWithSalt)

	return finalString
}

// Альтернативный вариант с использованием цикла
// func convert_to_rune_slice(s string) []rune {
// 	var runeSlice []rune
// 	for _, r := range s {
// 		runeSlice = append(runeSlice, r)
// 	}
// 	return runeSlice
// }
