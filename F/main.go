package main

import "fmt"

func main() {

	var (
		first  string
		second string
		third  string
		fourth string
		author string
	)

	fmt.Scan(&first, &second, &third, &fourth, &author)

	fmt.Printf("%s - %s\n", fourth, author)
	fmt.Printf("%s - %s\n", third, author)
	fmt.Printf("%s - %s\n", second, author)
	fmt.Printf("%s - %s\n", first, author)
}
