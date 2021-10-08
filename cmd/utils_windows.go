// +build windows

package cmd

import (
	"os/exec"
	"strings"
)

func OS() string {
	// use wmic for getting caption with OS name
	out, _ := exec.Command("wmic", "os", "get", "caption").Output()
	windowsVer := strings.TrimSpace(strings.Split(string(out), "\r\r\n")[1])

	return strings.TrimPrefix(windowsVer, "Microsoft ")
}

func Kernel() string {
	// use wmic for getting OS version
	out, _ := exec.Command("wmic", "os", "get", "Version").Output()
	kernelVer := strings.TrimSpace(strings.Split(string(out), "\r\r\n")[1])

	return kernelVer
}

func Shell() string {
	return "cmd.exe"
}

func WM() string {
	return "Metro"
}
