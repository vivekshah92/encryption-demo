package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	thirdparty "github.com/vivekshah92/encryption-demo/third_party"
)

func ListContentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "greet",
		Short: "Greet User",
		Long:  `Greet User`,
		Run: func(cmd *cobra.Command, args []string) {
			name := thirdparty.SetName("Encloud")
			fmt.Fprintf(cmd.OutOrStdout(), string(name))
		},
	}

	return cmd
}

func init() {
	RootCmd.AddCommand(ListContentsCmd())
}
