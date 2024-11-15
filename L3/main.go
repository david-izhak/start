package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		age  int
		name string
	)

	//fmt.Scan(&age)
	//fmt.Scan(&name)
	fmt.Scan(&age, &name)

	fmt.Printf("My name is: %s. My age is: %d\n", name, age)

	fmt.Fscan(os.Stdin, &age)
	fmt.Printf("New age is: %d\n", age)
}
