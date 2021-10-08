package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
  ^ \   /___\  /\ |
     |__|   |__| -
`

// elefetch fetches system info
func elefetch() {
	fmt.Printf(
		ele,
		UserInfo(),
		OS(),
		Kernel(),
		Shell(),
	)
	printColors()
}
