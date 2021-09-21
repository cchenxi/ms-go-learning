package main

import (
	"fmt"
	"strings"
)

type triangle struct {
	size int
}

type square struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

func (s square) perimeter() int {
	return s.size * 4
}

type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper((string(s)))
}

type coloredTriangle struct {
	triangle
	color string
}

func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 2
}

func main() {
	t := triangle{3}
	s := square{4}
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Perimeter (square):", s.perimeter())

	t1 := triangle{3}
	t1.doubleSize()
	fmt.Println("Size:", t1.size)
	fmt.Println("Perimeter:", t1.perimeter())

	str := upperstring("Learning Go")
	fmt.Println(str)
	fmt.Println(str.Upper())

	t2 := coloredTriangle{triangle{3}, "red"}
	fmt.Println("Size:", t2.size)
	fmt.Println("Perimeter(normal):", t2.triangle.perimeter())
	fmt.Println("Perimeter(colored):", t2.perimeter())
}
