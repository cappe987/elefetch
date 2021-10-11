package main

import (
	"fmt"
	"os"

	"github.com/burntcarrot/elefetch/cmd"
)


func main() {
	elefetch := cmd.EleFetch
	elefetch.Flags().StringVar(&cmd.AsciiPath, "source", "", "Path to custom ASCII file")
	elefetch.ParseFlags(os.Args)
	
	if err := elefetch.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
