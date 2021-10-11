// +build linux

package cmd

import (
	"os"
	"os/exec"
	"strings"
)

// fetch OS name
func OS() string {
	var osname []byte

	osname, _ = exec.Command("bash", "-c", "lsb_release -d | awk -F ' ' '{$1=\"\"; print $0}'").CombinedOutput()
	if strings.Contains(string(osname), "command not found") {
		if _, err := os.Stat("/etc/os-release"); os.IsNotExist(err) {
			// /etc/os-release contains variables. Source and echo the prettyname.
			osname, _ = exec.Command("bash", "-c", "source /etc/os-release; echo $PRETTY_NAME").CombinedOutput()
		} else {
			// The other files contain just the description/prettyname.
			osname, _ = exec.Command("bash", "-c", "cat /etc/*-release | head -1").CombinedOutput()
			if strings.Contains(string(osname), "No such file or directory") {
				return "" // Default to empty string.
			}
		}
	}

	return strings.TrimSpace(string(osname))
}

// fetch kernel info
func Kernel() string {
	procVer, _ := os.ReadFile("/proc/version")
	return strings.TrimSpace(strings.Split(string(procVer), " ")[2])
}

// fetch shell name
func Shell() string {
	shellEnv := strings.Split(os.Getenv("SHELL"), "/")
	return strings.TrimSpace(shellEnv[len(shellEnv)-1])
}

// fetch window manager name
func WM() string {
	wm, _ := exec.Command("bash", "-c", `xprop -id $(xprop -root -notype | awk '$1=="_NET_SUPPORTING_WM_CHECK:"{print $5}') -notype -f _NET_WM_NAME 8t | grep "WM_NAME" | cut -f2 -d \"`).CombinedOutput()
	return strings.TrimSpace(string(wm))
}
