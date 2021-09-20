package main

import (
	"fmt"
	"strconv"
)

func main() {
	// practice 1
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzBuzz(i))
	}

	// practice 2
	for number := 1; number < 20; number++ {
		if findprimes(number) {
			fmt.Printf("%v\n", number)
		}
	}

	// practice 3
	val := 0
	for {
		fmt.Println("enter number: ")
		fmt.Scanf("%d", &val)

		switch {
		case val < 0:
			panic("error")
		case val == 0:
			fmt.Println("0 is neither negative nor positive.")
		default:
			fmt.Println("you entered: ", val)
		}
	}
}

func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		if number % i == 0 {
			return false
		}
	}

	if number > 1 {
		return true
	} else {
		return false
	}
}

func fizzBuzz(num int) string {
	switch {
	case num %3 == 0 && num % 5 == 0:
		return "FizzBuzz"
	case num % 3 == 0:
		return "Fizz"
	case num % 5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(num)
	}
}
