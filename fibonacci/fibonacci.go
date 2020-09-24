package fibonacci
//Written by Chris Hakala when asked as part of an interview process.

import (
    "fmt"
    "log"
)

// Takes a slice of fibonacci numbers and returns a slice that includes the next number
func fibonacciSlice(prev []int) []int {
    s := len(prev)
    var next int
    if s > 1 {
        next = prev[s-1] + prev[(s-2)]
    } else if s == 1 {
        next = 1
    } else {
        next = 0
    }
    return append(prev, next)
}

// Returns the sequence of fibonacci numbers from f(0) to f(n)
func Sequence(count int) [] int {
    var fib []int
    for temp := 0; temp <= count; temp++ {
        fib = fibonacciSlice(fib)
    }
    return fib
}

// Prints the sequence of fibonacci numbers from f(0) to f(n)
func PrintSequence(count int) {
    fib := Sequence(count)
    if len(fib) < 1 {
        _, err := fmt.Println("Empty sequence")
        if err != nil {
            log.Fatal(err)
        }
    } else {
        s := fmt.Sprintf("Sequence: %d", fib[0])
        if len(fib) > 1 {
            for _, c := range fib[1:] {
                s = fmt.Sprintf("%s, %d", s, c)
            }
        }
        _, err := fmt.Println(s)
        if err != nil {
            log.Fatal(err)
        }
    }
}