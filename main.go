package main

import (
	"fmt"
	"os"

	"github.com/burntcarrot/elefetch/cmd"
)


func main() {
	elefetch := cmd.EleFetch
	elefetch.Flags().StringVar(&cmd.AsciiPath, "source", "", "Path to custom ASCII file")
	err := elefetch.ParseFlags(os.Args)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Parsing arguments failed. Using default settings")
	}
	
	if err := elefetch.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
