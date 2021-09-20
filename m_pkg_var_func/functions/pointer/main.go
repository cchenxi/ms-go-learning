package main

import "fmt"

func main() {
	firstName := "john"
	updateName(&firstName)
	fmt.Println(firstName)
}

func updateName(name *string) {
	*name = "david"
}
