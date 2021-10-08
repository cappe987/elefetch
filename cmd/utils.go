package cmd

import (
	"os"
	"os/user"
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
