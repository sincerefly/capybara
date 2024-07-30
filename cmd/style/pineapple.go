package style

import (
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/cmd/cmdutils"
	"github.com/sincerefly/capybara/service/style"
	"github.com/sincerefly/capybara/utils/colorizer"
	"github.com/spf13/cobra"
)

var PineappleCmd = &cobra.Command{
	Use:   "pineapple",
	Short: "Style: retro style film time",
	Run: func(cmd *cobra.Command, args []string) {

		parameter := &style.PineappleParameter{}

		input := cmdutils.GetParam(cmd.Flags(), "input")
		parameter.SetInput(input)

		output := cmdutils.GetParam(cmd.Flags(), "output")
		parameter.SetOutput(output)

		// color param
		colorStr := cmdutils.GetParam(cmd.Flags(), "color")
		col, err := colorizer.ToColor(colorStr)
		if err != nil {
			log.Fatal(err)
		}
		parameter.SetFontColor(col)

		// run
		log.Debugf("parameter: %s", parameter.JSONString())
		style.NewStyleProcessor(style.StylePineapple, parameter).Run()
	},
}

func init() {

	flags := PineappleCmd.Flags()
	flags.StringP("input", "i", "input", "specify input folder")
	flags.StringP("output", "o", "output", "specify output folder")
	flags.StringP("color", "c", "rgba(255, 140, 0, 230)", "specify font color")
}
