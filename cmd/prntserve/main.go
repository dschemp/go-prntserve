package main

import (
	"fmt"
	internal "github.com/dschemp/go-prntserve/internal"
	"github.com/dschemp/go-prntserve/internal/cmd"
	"github.com/dschemp/go-prntserve/internal/handler"
	"log"
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

	log.Printf(`Probing storage path "%s"...`, cmd.StoragePath())
	err := handler.ProbeStoragePathOnFS()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Probing storage path successfully.")

	// Start server
	internal.StartServer()
}
