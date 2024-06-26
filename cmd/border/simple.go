package border

import (
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/service/border"
	"github.com/sincerefly/capybara/service/border/styles"
	"github.com/sincerefly/capybara/utils/cobra_utils"
	"github.com/sincerefly/capybara/utils/colorizer"
	"github.com/spf13/cobra"
)

var SimpleCmd = &cobra.Command{
	Use:   "simple",
	Short: "Style: add a uniform-width border to the image.",
	Run: func(cmd *cobra.Command, args []string) {

		parameter := &styles.SimpleParameter{}

		input := cobra_utils.GetParam(cmd.Flags(), "input")
		parameter.SetInput(input)

		output := cobra_utils.GetParam(cmd.Flags(), "output")
		parameter.SetOutput(output)

		// width param
		width := cobra_utils.GetIntParam(cmd.Flags(), "width")
		if fixedWidth, fixed := border.FixedBorderWidth(width); fixed {
			log.Warn("border width fixed with %d", width)
			width = fixedWidth
		}
		parameter.SetBorderWidth(width)

		// color param
		colorStr := cobra_utils.GetParam(cmd.Flags(), "color")
		col, err := colorizer.ToColor(colorStr)
		if err != nil {
			log.Fatal(err)
		}
		parameter.SetBorderColor(col)

		// run
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
