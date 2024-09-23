package main

import (
	"fmt"
	"os"
)

const textOfBadExec = "Usage: go run . <task1 | task2>"

func main() {

	if len(os.Args) != 2 {
		fmt.Println(textOfBadExec)
		return
	}

	switch os.Args[1] {
	case "task1":
		var n int
		fmt.Print("Enter a number of iterations for task1: ")
		fmt.Scanln(&n)
		result := Task1(n)
		fmt.Println("Result from task2:", result)

	case "task2":
		var precision float64
		fmt.Print("Enter a precision for task2 (1, 0.1, 0.01 ...): ")
		fmt.Scanln(&precision)
		result := Task2(precision)
		fmt.Println("Result from task2", result)

	default:
		fmt.Println(textOfBadExec)
	}

}
