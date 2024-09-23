package main

import (
	"fmt"
	"os"
	"time"
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
		go spiner(100)
		result := Task1(n)
		fmt.Printf("\rtask1(%v): %v", n, result)

	case "task2":
		var precision float64
		fmt.Print("Enter a precision for task2 (1, 0.1, 0.01 ...): ")
		fmt.Scanln(&precision)
		go spiner(100)
		result := Task2(precision)
		fmt.Printf("\rtask2(%v): %v", precision, result)

	default:
		fmt.Println(textOfBadExec)
	}
}

func spiner(delay time.Duration) {
	for {
		for _, i := range `\|/-` {
			fmt.Printf("\r%c", i)
			time.Sleep(delay * time.Millisecond)
		}
	}
}
