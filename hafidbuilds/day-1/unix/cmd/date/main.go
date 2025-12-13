package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [+FORMAT]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Display the current local time and given FORMAT.\n\n")
		fmt.Fprintf(os.Stderr, "Supported:\n")
		fmt.Fprintf(os.Stderr, "  +\"%%Y-%%m-%%d\"    ISO date (2025-12-10)\n")
		fmt.Fprintf(os.Stderr, "  +\"%%H:%%M:%%S\"    Time (14:30:45)\n")
		fmt.Fprintf(os.Stderr, "\nIf no format is specified, displays default format.\n")
	}
	flag.Parse()
	now := time.Now()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println(now.Format("Mon Jan 2 15:04:05 MST 2006"))
		return
	}

	if !strings.HasPrefix(args[0], "+") {
		fmt.Fprintf(os.Stderr, "Error: format must start with +\n")
		os.Exit(1)
	}

	format := args[0][1:]
	output := formatDate(format, now)
	fmt.Println(output)
}

func formatDate(format string, t time.Time) string {
	switch format {
	case "%Y-%m-%d":
		return t.Format("2006-01-02")
	case "%H:%M:%S":
		return t.Format("15:04:05")
	default:
		fmt.Fprintf(os.Stderr, "Error: unsupported format\n")
		os.Exit(1)
		return ""
	}
}
