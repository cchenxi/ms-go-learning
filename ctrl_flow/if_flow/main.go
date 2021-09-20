package main

import "fmt"

func main() {
	x := 27
	if x % 2 == 0 {
		fmt.Println(x, " is even.")
	}

	if num := giveMeanNumber(); num < 0 {
		fmt.Println(num, " is negative.")
	} else if num < 10 {
		fmt.Println(num, "has only one digit.")
	} else {
		fmt.Println(num, " has multi digits.")
	}
}

func giveMeanNumber() int {
	return -1
}
