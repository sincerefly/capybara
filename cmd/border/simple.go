package border

import (
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/cmd/border_common"
	"github.com/sincerefly/capybara/cmd/cmdutils"
	"github.com/sincerefly/capybara/service/border"
	"github.com/sincerefly/capybara/service/border/styles"
	"github.com/sincerefly/capybara/utils/colorizer"
	"github.com/spf13/cobra"
)

var SimpleCmd = &cobra.Command{
	Use:   "simple",
	Short: "Style: add a uniform-width border to the image.",
	Run: func(cmd *cobra.Command, args []string) {

		parameter := &styles.SimpleParameter{}

		input := cmdutils.GetParam(cmd.Flags(), "input")
		parameter.SetInput(input)

		output := cmdutils.GetParam(cmd.Flags(), "output")
		parameter.SetOutput(output)

		// width param
		width := cmdutils.GetIntParam(cmd.Flags(), "width")
		if fixedWidth, fixed := border_common.FixedBorderWidth(width); fixed {
			log.Warn("border width fixed with %d", width)
			width = fixedWidth
		}
		parameter.SetBorderWidth(width)

		// color param
		colorStr := cmdutils.GetParam(cmd.Flags(), "color")
		col, err := colorizer.ToColor(colorStr)
		if err != nil {
			log.Fatal(err)
		}
		parameter.SetBorderColor(col)

		// run
		log.Debugf("parameter: %s", parameter.JSONString())
		border.NewStyleProcessor(border.StyleSimple, parameter).Run()
	},
}

func init() {

	flags := SimpleCmd.Flags()
	flags.StringP("input", "i", "input", "specify input folder")
	flags.StringP("output", "o", "output", "specify output folder")
	flags.IntP("width", "w", 100, "specify border width")
	flags.StringP("color", "c", "white", "specify border color")
}
