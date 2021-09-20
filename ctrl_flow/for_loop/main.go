package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println("sum of 1..100 is ", sum)

	// 模拟while 循环
	whileLoop()
	// 模拟死循环
	deadLoop()
	// continue 关键字
	keyContinue()
}

func whileLoop() {
	var num int64
	rand.Seed(time.Now().Unix())
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Println(num)
	}
}

func deadLoop() {
	var num int32
	sec := time.Now().Unix()
	rand.Seed(sec)

	for {
		fmt.Println("writing inside the loop ...")
		if num = rand.Int31n(10); num == 5 {
			fmt.Println("finish!")
			break
		}
		fmt.Println(num)
	}
}

func keyContinue() {
	sum := 0
	for num := 1; num <= 100; num++ {
		if num % 5 == 0 {
			continue
		}
		sum += num
	}

	fmt.Println("sum of 1 to 100, but excluding numbers divisible by 5, is ", sum)
}


