package maker

import (
	"github.com/spf13/cobra"
)

var makerMakeAdminModuleCmd = &cobra.Command{
	Use:   "admin-module [name]",
	Short: "Make an admin module",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		/** @todo implement command */

		return nil
	},
}

func init() {
	extensionMakerRootCommand.AddCommand(makerMakeAdminModuleCmd)
	makerMakeAdminModuleCmd.PersistentFlags().String("name", "", "name of the module")
}
