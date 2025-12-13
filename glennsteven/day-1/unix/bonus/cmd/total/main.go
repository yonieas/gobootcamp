package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := 0

	for scanner.Scan() {
		text := scanner.Text()
		m, err := strconv.Atoi(text)
		if err != nil {
			continue
		}

		total += m
	}

	fmt.Println(total)
}
