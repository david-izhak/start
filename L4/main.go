package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

type Person struct {
	firstname string
	lastname  string
	age       int8
}

func main() {
	person := Person{"Иван", "Иванов", 25}
	fmt.Printf("Имя: %s, Фамилия: %s\nРазмер объекта: %d byte\n", person.firstname, person.lastname, unsafe.Sizeof(person))

	var num1 int64 = 1
	var num2 uint64 = 2
	var num3 float64 = 3.7

	fmt.Println(num1 + int64(num2))
	fmt.Println(num1 + int64(num3))

	str1 := "Hello"
	fmt.Println("Length of string in bytes:", len(str1))
	fmt.Println("Length of string in runes:", len([]rune(str1)))
	fmt.Println("Length of string in runes (chars):", utf8.RuneCountInString(str1)) // utf8.RuneCountInString - возвращает количество рун в строке. Rune - это один символ в UTF

	str2 := "Привет"
	fmt.Println("Length of string in bytes:", len(str2))
	fmt.Println("Length of string in runes:", len([]rune(str2)))
	fmt.Println("Length of string in runes (chars):", utf8.RuneCountInString(str2))

	fullString, substring := "Hello, World!", "World"
	fmt.Println(strings.Contains(fullString, substring))

	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("b", "a"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println("a" > "b")
	fmt.Println("b" > "a")
	fmt.Println("a" == "a")
	fmt.Println("a" > "ab")
	fmt.Println("a" > "ba")

}
