package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: seq <start> <end>")
		return
	}

	start, _ := strconv.Atoi(os.Args[1])
	end, _ := strconv.Atoi(os.Args[2])

	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
}
