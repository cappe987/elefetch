package cmd

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

// UserInfo fetches basic information about the user.
func UserInfo() string {
	// fetch current user
	currentUser, err := user.Current()
	username := "Unknown"

	if err == nil {
		// set username
		username = currentUser.Username
	}

	// fetch hostname
	hostname, err := os.Hostname()
	host := "Unknown"

	if err == nil {
		host = hostname
	}

	return username + "@" + host
}

func printColors(paddingLength int) {
	// whitespace for padding
	padding := strings.Repeat(" ", paddingLength)

	fmt.Printf("\n" + padding)
	for i := 0; i < 8; i++ {
		fmt.Printf("\033[4%dm   ", i)
	}
	fmt.Printf("\033[0m" + "\n" + padding)

	for i := 0; i < 8; i++ {
		fmt.Printf("\033[10%dm   ", i)
	}
	fmt.Printf("\033[0m" + "\n")
}
