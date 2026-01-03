package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		trimmedCommand := strings.TrimSpace(command)
		if trimmedCommand == "" {
			continue
		}

		params := strings.Fields(trimmedCommand)
		if len(params) == 0 {
			continue
		}

		switch params[0] {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Join(params[1:], " "))
		case "type":
			if len(params) != 2 {
				fmt.Fprintln(os.Stderr, "Invalid param passed")
				os.Exit(1)
			}
			param := params[1]
			if param == "exit" || param == "echo" || param == "type" {
				fmt.Printf("%s is a shell builtin\n", param)
			} else {
				fmt.Printf("%s: not found\n", param)
			}
		default:
			fmt.Printf("%s: command not found\n", trimmedCommand)
		}
	}
}
