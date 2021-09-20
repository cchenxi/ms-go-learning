package main

import (
	"fmt"
	"github.com/cchenxi/calculator"
)

func main() {
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println(calculator.Version)
}
