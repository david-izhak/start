package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		firstname string
		surname   string
		age       int
	)

	//fmt.Scan(&age)
	//fmt.Scan(&name)
	fmt.Scan(&firstname, &surname, &age)

	fmt.Fscan(os.Stdin, &age) // Позволяет считывать не только с консоли, но из других источников, например из файла. В данном примере считывает из консоли (os.Stdin)

	fmt.Printf("Имя: %s , Фамилия: %s , Возраст: %d . Студент BPS\n", firstname, surname, age)
}
