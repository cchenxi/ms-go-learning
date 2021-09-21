package main

import "fmt"

func main() {
	var num int
	fmt.Println("What's the Fibonacci sequence you want? ")
	_, _ = fmt.Scanln(&num)
	fmt.Println("The Fibonacci sequence is: ", Fibonacci(num))

	fmt.Println("MCLX is", romanToArabic("MCLX"))
	fmt.Println("MCMXCIX is ", romanToArabic("MCMXCIX"))
	fmt.Println("MCMZ is", romanToArabic("MCMZ"))
}

func Fibonacci(num int) []int {
	if num < 2 {
		return make([]int, 0)
	}
	nums := make([]int, num)
	nums[0], nums[1] = 1, 1
	for i := 2; i < num; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums
}

func romanToArabic(numeral string) int {
	romanMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	arabicVals := make([]int, len(numeral)+1)
	for index, digit := range numeral {
		if val, present := romanMap[digit]; present {
			arabicVals[index] = val
		} else {
			fmt.Println("Error")
			return 0
		}
	}

	total := 0

	for index := 0; index < len(numeral); index++ {
		if arabicVals[index] < arabicVals[index+1] {
			arabicVals[index] = -arabicVals[index]
		}
		total += arabicVals[index]
	}

	return total
}
