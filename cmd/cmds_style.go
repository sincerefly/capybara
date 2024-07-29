package cmd

import (
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/cmd/style"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(styleCmd)

	styleCmd.AddCommand(style.SimpleCmd)
	styleCmd.AddCommand(style.TextBottomCmd)
	styleCmd.AddCommand(style.LogoMelonCmd)
	styleCmd.AddCommand(style.PineappleCmd)
}

var styleCmd = &cobra.Command{
	Use:   "style",
	Short: "To batch add borders to images.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			log.Warnf("style '%s' unsupported", args[0])
		} else {
			log.Warnf("style need subcommand, e.g., 'capybara style simple [parameter]'")
		}
	},
}
