package main

import "fmt"

func main() {
	var a [3]int

	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a) - 1])

	//cities := [5]string{"new york", "shanghai", "nanjing", "beijing"}
	cities := [...]string{"shanghai", "nanjing", "beijing"}
	fmt.Println("Cities: ", cities)

	q := [...]int{1, 2, 3}
	fmt.Println(q)

	numbers := [...]int{99:-1}
	fmt.Println(numbers)
	fmt.Println("First position: ", numbers[0])
	fmt.Println("Last position: ", numbers[99])
	fmt.Println("Length: ", len(numbers))

	var twoDArr [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoDArr[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("row", i, twoDArr[i])
	}
	fmt.Println("\nAll at once: ", twoDArr)
}
