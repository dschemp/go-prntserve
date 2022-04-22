package main

import (
	"fmt"
	internal "github.com/dschemp/go-prntserve/internal"
	"github.com/dschemp/go-prntserve/internal/cmd"
	"os"
)

func main() {
	cmd.ParseArguments()
	fmt.Println("Welcome to prntserve!")

	if cmd.ShowVersion() {
		// TODO: Show version
		fmt.Printf("Version: %s\n", "unknown")
		os.Exit(0)
	}

	// Start server
	internal.StartServer()
}
