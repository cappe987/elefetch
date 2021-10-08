package cmd

import (
	"os"
	"os/exec"
	"strings"
)

// fetch OS name
func OS() string {
	return os.Getenv("NAME")
}

// fetch kernel info
func Kernel() string {
	procVer, _ := os.ReadFile("/proc/version")
	return strings.Split(string(procVer), " ")[2]
}

// fetch shell name
func Shell() string {
	shellEnv := strings.Split(os.Getenv("SHELL"), "/")
	return shellEnv[len(shellEnv)-1]
}

// fetch window manager name
func WM() string {
	wm, _ := exec.Command("bash", "-c", `xprop -id $(xprop -root -notype | awk '$1=="_NET_SUPPORTING_WM_CHECK:"{print $5}') -notype -f _NET_WM_NAME 8t | grep "WM_NAME" | cut -f2 -d \"`).CombinedOutput()
	return string(wm)
}
