package maker

import (
	"github.com/spf13/cobra"
)

var extensionMakerRootCommand = &cobra.Command{
	Use:   "make",
	Short: "Maker for boilerplate code in a shopware extension",
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(extensionMakerRootCommand)
}
