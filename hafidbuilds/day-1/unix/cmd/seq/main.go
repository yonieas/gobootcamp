package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func parseArg(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "seq: invalid argument: %s\n", s)
		os.Exit(1)
	}
	return val
}

func printNumber(n float64) {
	if n == float64(int(n)) {
		fmt.Printf("%d\n", int(n))
	} else {
		fmt.Printf("%g\n", n)
	}
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... LAST\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "   or: %s [OPTION]... FIRST LAST\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "   or: %s [OPTION]... FIRST INCREMENT LAST\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Print numbers from FIRST to LAST, in steps of INCREMENT.\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 || len(args) > 3 {
		flag.Usage()
		os.Exit(1)
	}

	var first, increment, last float64

	switch len(args) {
	case 1:
		first = 1
		increment = 1
		last = parseArg(args[0])
	case 2:
		first = parseArg(args[0])
		increment = 1
		last = parseArg(args[1])
	case 3:
		first = parseArg(args[0])
		increment = parseArg(args[1])
		last = parseArg(args[2])
	}

	if increment == 0 {
		fmt.Fprintf(os.Stderr, "seq: INCREMENT can't be 0\n")
		os.Exit(1)
	}

	if increment > 0 {
		for i := first; i <= last; i += increment {
			printNumber(i)
		}
	} else {
		for i := first; i >= last; i += increment {
			printNumber(i)
		}
	}
}
