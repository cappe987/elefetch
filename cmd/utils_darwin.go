// +build darwin

package cmd

import (
	"os"
	"os/exec"
	"strings"
)

func OS() string {
	out, err := exec.Command("sw_vers", "-productName").CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	return strings.TrimSuffix(string(out), "\n")
}

func Kernel() string {
	out, err := exec.Command("uname", "-r").CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	return strings.TrimSuffix(string(out), "\n")
}

func Shell() string {
	shellEnv := strings.Split(os.Getenv("SHELL"), "/")
	return shellEnv[len(shellEnv)-1]
}

func WM() string {
	return "Quartz"
}
