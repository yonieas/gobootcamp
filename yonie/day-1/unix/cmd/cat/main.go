package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		cat(os.Stdin)
		return
	}

	// Iterate over provided files
	for i := 1; i < len(os.Args); i++ {
		filename := os.Args[i]
		// Open file
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %s\n", filename, err)
			// Skip if failed
			continue
		}
		// Process file
		cat(f)
		// Close file
		f.Close()
	}
}

func cat(f *os.File) {
	if _, err := io.Copy(os.Stdout, f); err != nil {
		fmt.Fprintf(os.Stderr, "cat: error reading from %s: %s\n", f.Name(), err)
		os.Exit(1)
	}
}
