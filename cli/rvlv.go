package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		usageMenu()
		os.Exit(0)
	}

	switch args[0] {
	case "init":
		fmt.Println("init")
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
	os.Exit(0)
}

func usageMenu() {
	fmt.Println("Usage: rvlv init")
}
