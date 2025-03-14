package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"
)

const (
	BackgroundRed    = "\033[41m"
	BackgroundGreen  = "\033[42m"
	BackgroundYellow = "\033[43m"
	BackgroundBlue   = "\033[44m"
	Reset            = "\033[0m"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {

		now := time.Now().Format("02 Jan 06 15:04")

		fmt.Print("🕑", now)

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Println("", cwd, ">   origin ☊ master 2☀")

		// Get the current user information.
		user, err := user.Current()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Print("@", user.Username, ">")

		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// ErrNoPath is returned when 'cd' was called without a second argument.
var ErrNoPath = errors.New("path required")

func execInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input separate the command and the arguments.
	args := strings.Split(input, " ")

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		// 'cd' to home with empty path not yet supported.
		if len(args) < 2 {
			return ErrNoPath
		}
		// Change the directory and return the error.
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}
