package cmd

import (
	"fmt"

	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Shows a short information about Composer.",
	Long:  `Shows a short information about Composer.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(color.Colorize(color.Green, "Melodist - Dependency Manager for PHP"))
		fmt.Println(color.Colorize(color.Yellow, "Melodist is a dependency manager tracking local dependencies of your projects and libraries."))
	},
}
