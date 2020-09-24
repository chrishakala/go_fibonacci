package main
//Written to demonstrate fibonacci package.

import (
	"fmt"
	"os"
	"strconv"
	"../fibonacci"
)

func main() {
    count := 20
	if len(os.Args) > 1 {
	    temp, err := strconv.Atoi(os.Args[1])
		if err != nil {
		    fmt.Println("Invalid command line Argument. defaulting to 20")
		} else {
			count = temp
		}
		fmt.Printf("Reporting fibonacci sequence from f(0) to f(%d)\r\n", count)
	} else {
		fmt.Println("Reporting fibonacci sequence from f(0) to f(20)")
	}
    fibonacci.PrintSequence(count)
}