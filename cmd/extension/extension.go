package extension

import (
	"github.com/FriendsOfShopware/shopware-cli/cmd/extension/maker"
	"github.com/spf13/cobra"
)

var extensionRootCmd = &cobra.Command{
	Use:   "extension",
	Short: "Shopware Extension utilities",
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(extensionRootCmd)
	maker.Register(extensionRootCmd)
}
