package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Handle if the parameter more than 2
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <start_number> <end_number>\n", os.Args[0])
		return
	}
	s, err := isNumber(os.Args[1])
	if !err {
		return
	}
	e, err := isNumber(os.Args[2])
	if !err {
		return
	}
	GenerateNumbers(s, e)
}

func GenerateNumbers(start int, end int) {
	if start > end {
		for i := start; i >= end; i-- {
			fmt.Println(i)
		}
	} else {
		for i := start; i <= end; i++ {
			fmt.Println(i)
		}
	}
}

func isNumber(s string) (int, bool) {
	num, errInt := strconv.Atoi(s)
	if errInt != nil {
		fmt.Fprintf(os.Stderr, "Parameter '%s' is not a valid number\n", s)
		return 0, false
	}
	return num, true
}
