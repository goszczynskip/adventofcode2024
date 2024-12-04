package main

import (
	"adventofcode/internal/app"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command line flags
	day := flag.String("day", "", "Day to run eg. 1_1")
	test := flag.Bool("test", false, "Use test input")
	flag.Parse()

	if err := app.Run(*day, *test); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
