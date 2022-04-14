package cmd

import (
	"github.com/rsteube/carapace"
	bat "github.com/rsteube/carapace-bin/completers/bat_completer/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "prettybat",
	Short:              "Pretty-print source code and highlight it with bat",
	Long:               "https://github.com/eth-p/bat-extras/blob/master/doc/prettybat.md",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bat.ActionExecute(),
	)
}
