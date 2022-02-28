package cmd

import (
	"fmt"

	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

type aboutCmd struct {
	cmd *cobra.Command
}

func (c *aboutCmd) getCommand() *cobra.Command {
	return c.cmd
}

func newAboutCmd() Command {
	cc := &aboutCmd{}

	cc.cmd = &cobra.Command{
		Use:   "about",
		Short: "Shows a short information about Composer.",
		Long:  `Shows a short information about Composer.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(color.Colorize(color.Green, "Melodist - Dependency Manager for PHP"))
			fmt.Println(color.Colorize(color.Yellow, "Melodist is a dependency manager tracking local dependencies of your projects and libraries."))
		},
	}

	return cc
}
