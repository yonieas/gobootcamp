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
		fmt.Println("Usage: filter <file>")
		return
	}

	mode := strings.ToLower(os.Args[1])
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		m, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}

		//'odd' refers to odd numbers (bilangan ganjil)
		if mode == "odd" && m%2 == 1 {
			fmt.Println(m)
		}
		//'even' refers to even numbers (bilangan genap)
		if mode == "even" && m%2 == 0 {
			fmt.Println(m)
		}
	}
}
