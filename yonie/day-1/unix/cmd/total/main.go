package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	calculateTotal()
}

func calculateTotal() {
	scanner := bufio.NewScanner(os.Stdin)

	total := 0
	for scanner.Scan() {
		numString := scanner.Text()
		num, err := strconv.Atoi(numString)
		if err == nil {
			total += num
		}
	}
	fmt.Fprintln(os.Stdout, total)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
	}
}
