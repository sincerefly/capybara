package border

import (
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/service/border"
	"github.com/sincerefly/capybara/service/border/styles"
	"github.com/sincerefly/capybara/utils/cobra_utils"
	"github.com/sincerefly/capybara/utils/colorizer"
	"github.com/spf13/cobra"
)

var TextBottomCmd = &cobra.Command{
	Use:   "text_bottom",
	Short: "Style: Footer text, with photo exif",
	Run: func(cmd *cobra.Command, args []string) {

		parameter := &styles.TextBottomParameter{}

		input := cobra_utils.GetParam(cmd.Flags(), "input")
		parameter.SetInput(input)

		output := cobra_utils.GetParam(cmd.Flags(), "output")
		parameter.SetOutput(output)

		// width param
		width := cobra_utils.GetIntParam(cmd.Flags(), "width")
		if fixedWidth, fixed := border.FixedBorderWidth(width); fixed {
			log.Warn("border width fixed with %d", fixedWidth)
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

		// bottom container height
		containerHeight := cobra_utils.GetIntParam(cmd.Flags(), "container-height")
		parameter.SetBottomContainerHeight(containerHeight)

		// with subtitle
		withoutSubtitle := cobra_utils.GetBoolParam(cmd.Flags(), "without-subtitle")
		parameter.SetWithoutSubtitle(withoutSubtitle)

		// run
		log.Debugf("parameter: %s", parameter.JSONString())
		border.NewStyleProcessor(border.StyleTextBottom, parameter).Run()
	},
}

func init() {

	flags := TextBottomCmd.Flags()
	flags.StringP("input", "i", "input", "specify input folder")
	flags.StringP("output", "o", "output", "specify output folder")
	flags.IntP("width", "w", 100, "specify border width")
	flags.StringP("color", "c", "white", "specify border color")
	flags.IntP("container-height", "", 300, "bottom text container height")
	flags.BoolP("without-subtitle", "", false, "without subtitle")
}
