package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num1, num2 := os.Args[1], os.Args[2]

	sum, mul := calc(num1, num2)
	fmt.Printf("sum = %d, mul = %d\n", sum, mul)

	_, mul2 := calc(num1, num2)
	fmt.Printf("mul = %d\n", mul2)
}

func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2
}

func calc(number1 string, number2 string) (int, int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2, int1 - int2
}
