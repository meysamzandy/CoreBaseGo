package main

import (
	"CoreBaseGo/internal/interfaces/cli/commands"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "corebase",
	Short: "CoreBase CLI is a tool for managing your CoreBaseGo application",
	Run: func(cmd *cobra.Command, args []string) {
		// Root command logic, executed when no subcommand is specified
		fmt.Println("Welcome to CoreBase CLI!")
	},
}

func init() {
	// Add subcommands to the root command
	rootCmd.AddCommand(commands.GreetCmd)
}

func main() {
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Println("Error executing CLI command:", err)
		os.Exit(1)
	}
}
