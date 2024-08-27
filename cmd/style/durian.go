package style

import (
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/cmd/border_common"
	"github.com/sincerefly/capybara/cmd/cmdutils"
	"github.com/sincerefly/capybara/service/style"
	"github.com/spf13/cobra"
)

var DurianCmd = &cobra.Command{
	Use:   "durian",
	Short: "Style: Gaussian blur background",
	Run: func(cmd *cobra.Command, args []string) {

		parameter := &style.DurianParameter{}

		input := cmdutils.GetParam(cmd.Flags(), "input")
		parameter.SetInput(input)

		output := cmdutils.GetParam(cmd.Flags(), "output")
		parameter.SetOutput(output)

		// width param
		width := cmdutils.GetIntParam(cmd.Flags(), "width")
		if fixedWidth, fixed := border_common.FixedBorderWidth(width); fixed {
			log.Warn("border width fixed with %d", fixedWidth)
			width = fixedWidth
		}
		parameter.SetBorderWidth(width)

		// with subtitle
		withoutSubtitle := cmdutils.GetBoolParam(cmd.Flags(), "without-subtitle")
		parameter.SetWithoutSubtitle(withoutSubtitle)

		// run
		log.Debugf("parameter: %s", parameter.JSONString())
		style.NewStyleProcessor(style.StyleDurian, parameter).Run()
	},
}

func init() {

	flags := DurianCmd.Flags()
	flags.StringP("input", "i", "input", "specify input folder")
	flags.StringP("output", "o", "output", "specify output folder")
	flags.IntP("width", "w", 500, "specify border width")
	flags.BoolP("without-subtitle", "", false, "without subtitle")
}
