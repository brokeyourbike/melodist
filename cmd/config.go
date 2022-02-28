package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

type configCmd struct {
	cmd    *cobra.Command
	global bool
	list   bool
	auth   string
	file   string
}

func (c *configCmd) getCommand() *cobra.Command {
	return c.cmd
}

func newConfigCmd() Command {
	cc := &configCmd{}

	cmd := &cobra.Command{
		Use:   "config",
		Short: "Sets config options.",
		Long: `This command allows you to edit composer config settings and repositories
in either the local composer.json file or the global config.json file.

Additionally it lets you edit most properties in the local composer.json.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		RunE: func(cmd *cobra.Command, args []string) error {
			if cc.global && cc.file != "" {
				return errors.New("file and global can not be combined")
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&cc.global, "global", "g", false, "Apply command to the global config file")
	cmd.Flags().BoolVarP(&cc.list, "list", "l", false, "List configuration settings")
	cmd.Flags().StringVarP(&cc.auth, "auth", "a", "", "Affect auth config file (only used for --editor)")
	cmd.Flags().StringVarP(&cc.file, "file", "f", "", "If you want to choose a different composer.json or config.json")

	cc.cmd = cmd

	return cc
}
