package maker

import (
	"github.com/spf13/cobra"
)

var makerMakeAdminComponentCmd = &cobra.Command{
	Use:   "admin-component [name] [moduleName]",
	Short: "Make an admin component inside an admin module",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		/** @todo implement command */

		return nil
	},
}

func init() {
	extensionMakerRootCommand.AddCommand(makerMakeAdminComponentCmd)
	makerMakeAdminComponentCmd.PersistentFlags().String("name", "", "name of the component")
	makerMakeAdminComponentCmd.PersistentFlags().String("module", "", "name of the module where the component should be created")
}
