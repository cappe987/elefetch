package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var AsciiPath string

// elefetch command
var EleFetch = &cobra.Command{
	Use:   "elefetch",
	Short: "Cross-platform system utility for fetching system info.",
	Run: func(cmd *cobra.Command, args []string) {
		elefetch()
	},
}

// elephant ASCII art
var ele = `
            __     __
           /  \---/  \  %s
     ,----(     ..    ) %s
    /      \__     __/  %s
   /|         (\  |(    %s
  ^ \   /___\  /\ |     %s
     |__|   |__| -
`

func customAscii(path string){
	values := [5]string{UserText(), OSText(), KernelText(), ShellText(), WMText()}
	path = strings.Replace(path, "~", os.Getenv("HOME"), 1) // Handles paths using ~/ syntax
	art, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error: no such file '%s'", path)
		return
	}

	lines := strings.Split(string(art), "\n")
	lines = lines[:len(lines)-1] // Cut off the extra empty line that is always there

	padLength := len(lines[0])
	fill := strings.Repeat(" ", padLength)
	// Add extra lines if the art is fewer lines than data
	for i := len(lines); i < 5; i++ {
		lines = append(lines, fill)
	}

	// Format data strings
	for i := 0; i < 5; i++ {
		lines[i] = fmt.Sprintf("%s %s", lines[i], values[i])
	}

	lines = append(lines, "") 
	fmt.Print(strings.Join(lines, "\n"))
	printColors(padLength + 2)
}

// elefetch fetches system info
func elefetch() {
	if AsciiPath != "" {
		customAscii(AsciiPath)
		return
	}
	fmt.Printf(
		ele,
		UserText(),
		OSText(),
		KernelText(),
		ShellText(),
		WMText(),
	)
	printColors(25)
}
