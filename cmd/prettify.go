package cmd

import "os"

var esc = "\033[0m"

func UserText() string {
	return "\033[31m" + UserInfo() + esc
}

func OSText() string {
	// get color from OS, if specified
	color := os.Getenv("ANSI_COLOR")

	if color == "" {
		color = "32"
	}

	return "\033[" + color + "m" + "OS: " + esc + OS()
}

func KernelText() string {
	return "\033[33m" + "Kernel: " + esc + Kernel()
}

func ShellText() string {
	return "\033[34m" + "Shell: " + esc + Shell()
}

func WMText() string {
	return "\033[35m" + "WM: " + esc + WM()
}
