package main

import (
	"fmt"
	"os"

	"github.com/burntcarrot/elefetch/cmd"
)

func main() {
	if err := cmd.EleFetch.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
