package cmd

import (
	"github.com/sincerefly/capybara/cmd/border"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(borderCmd)

	borderCmd.AddCommand(border.SimpleCmd)
	borderCmd.AddCommand(border.TextBottomCmd)
}

var borderCmd = &cobra.Command{
	Use:   "border",
	Short: "To batch add borders to images.",
	Run:   func(cmd *cobra.Command, args []string) {},
}
