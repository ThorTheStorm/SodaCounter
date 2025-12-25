package cli

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func PresentOptions(options []string) {
	fmt.Println("What do you want to do?")
	fmt.Println("-----------------------")
	for index, value := range options {
		fmt.Printf("%d: %s\n", index+1, value)
	}
}

// ClearScreen clears the terminal screen based on the operating system
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Windows command
	} else {
		cmd = exec.Command("clear") // Unix-based command
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
