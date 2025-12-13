package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <odd|even>\n", os.Args[0])
		os.Exit(1)
	}

	filterType := os.Args[1]

	if filterType != "odd" && filterType != "even" {
		fmt.Fprintf(os.Stderr, "Error: filter criteria must be 'odd' or 'even'\n")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		num, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		if filterType == "odd" && num%2 != 0 {
			fmt.Println(num)
		} else if filterType == "even" && num%2 == 0 {
			fmt.Println(num)
		}
	}
}
