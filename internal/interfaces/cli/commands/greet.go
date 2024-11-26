package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GreetCmd = &cobra.Command{
	Use:   "greet [name]",
	Short: "Greet a user with a name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Hello, %s! Welcome to CoreBaseGo!\n", name)
	},
}
