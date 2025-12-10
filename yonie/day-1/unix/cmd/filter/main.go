package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FilterType int

const (
	NoFilter FilterType = iota
	OddFilter
	EvenFilter
)

var filterName = map[FilterType]string{
	NoFilter:   "",
	OddFilter:  "odd",
	EvenFilter: "even",
}

func (ft FilterType) String() string {
	return filterName[ft]
}

func main() {
	// Check argument
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s 'odd' or 'even'\n", os.Args[0])
		os.Exit(1)
	}
	arg := os.Args[1]

	// convert string to enum
	filterType := convertStringFilter(arg)

	// filter using enum
	filter(filterType)
}

func convertStringFilter(s string) FilterType {
	var filterType FilterType = NoFilter
	switch strings.ToLower(s) {
	case OddFilter.String():
		filterType = OddFilter
	case EvenFilter.String():
		filterType = EvenFilter
	default:
		fmt.Fprintf(os.Stderr, "Error: Invalid Filter '%s'. Must be '%s' or '%s'.\n ", s, OddFilter.String(), EvenFilter.String())
		os.Exit(1)
	}
	return filterType
}

func filter(ft FilterType) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		numString := scanner.Text()
		num, err := strconv.Atoi(numString)
		if err == nil {
			if ft == OddFilter && num%2 != 0 {
				fmt.Fprintf(os.Stdout, "%d\n", num)
			} else if ft == EvenFilter && num%2 == 0 {
				fmt.Fprintf(os.Stdout, "%d\n", num)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
	}
}
