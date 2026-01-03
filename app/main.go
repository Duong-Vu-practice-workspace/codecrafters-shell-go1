package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print
var builtins = map[string]struct{}{
	"exit": {},
	"echo": {},
	"type": {},
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println()
				return
			}
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		fields := strings.Fields(command)
		if len(fields) == 0 {
			continue
		}
		command = fields[0]
		args := fields[1:]

		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Join(args, " "))
		case "type":
			if len(args) == 0 {
				fmt.Fprintln(os.Stderr, "type: missing operand")
				continue
			}
			cmdLookup := args[0]
			if _, ok := builtins[cmdLookup]; ok {
				fmt.Printf("%s is a shell builtin\n", cmdLookup)
			} else {
				fmt.Printf("%s: not found\n", cmdLookup)
			}
		default:
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
