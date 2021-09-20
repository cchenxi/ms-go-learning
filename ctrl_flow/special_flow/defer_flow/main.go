package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for i := 1; i <= 4; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}

	newFile, error := os.Create("learnGo.txt")
	if error != nil {
		fmt.Println("Error: Could not create file.")
		return
	}

	defer newFile.Close()

	if _, error = io.WriteString(newFile, "Learning Golang!"); error != nil {
		fmt.Println("Error: Could not write to file.")
	}

	newFile.Sync()
}
