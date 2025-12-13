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
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}

		total += num
	}

	fmt.Println(total)
}
