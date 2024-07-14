package cmd

import (
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/cmd/border"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(borderCmd)

	borderCmd.AddCommand(border.SimpleCmd)
	borderCmd.AddCommand(border.TextBottomCmd)
	borderCmd.AddCommand(border.LogoMelonCmd)
}

var borderCmd = &cobra.Command{
	Use:   "border",
	Short: "To batch add borders to images.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			log.Warnf("border style '%s' unsupported", args[0])
		} else {
			log.Warnf("border need subcommand, e.g., 'capybara border simple [parameter]'")
		}
	},
}
