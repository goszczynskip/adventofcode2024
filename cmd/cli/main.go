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
	debug := flag.Bool("debug", false, "Debug run")
	flag.Parse()

	if *debug {
		fmt.Println("Debug mode")

		// Run the application
		if err := app.Debug(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	} else if err := app.Run(*day, *test); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
