package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [FILE]...\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Concat FILES to standard output.\n\n")
		fmt.Fprintf(os.Stderr, "With no FILE, or when FILE is -, read standard input.\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		catFile(os.Stdin)
		return
	}

	exitCode := 0
	for _, filename := range args {
		if filename == "-" {
			catFile(os.Stdin)
		} else {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cat: %s: %v\n", filename, err)
				exitCode = 1
				continue
			}

			catFile(file)
			file.Close()
		}
	}

	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

func catFile(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
