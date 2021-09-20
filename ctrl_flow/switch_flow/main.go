package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)
	i = 3

	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
	case 4:
		fallthrough
	case 5:
		fmt.Println("Four, Five")
	default:
		fmt.Println(i)
	}

	fmt.Println("Ok")
}
