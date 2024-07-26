package style

import (
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/cmd/border_common"
	"github.com/sincerefly/capybara/cmd/cmdutils"
	"github.com/sincerefly/capybara/service/style"
	"github.com/sincerefly/capybara/utils/colorizer"
	"github.com/spf13/cobra"
)

var LogoMelonCmd = &cobra.Command{
	Use:   "melon",
	Short: "Style: logo left, no padding",
	Run: func(cmd *cobra.Command, args []string) {

		parameter := &style.LogoMelonParameter{}

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

		// color param
		colorStr := cmdutils.GetParam(cmd.Flags(), "color")
		col, err := colorizer.ToColor(colorStr)
		if err != nil {
			log.Fatal(err)
		}
		parameter.SetBorderColor(col)

		// bottom container height
		containerHeight, set := cmdutils.GetIntParamB(cmd.Flags(), "container-height")
		parameter.SetBottomContainerHeight(containerHeight)
		parameter.SetIsContainerHeightSet(set)

		// bottom container height
		containerHeightRatio := cmdutils.GetFloat64Param(cmd.Flags(), "container-height-ratio")
		parameter.SetBottomContainerHeightRatio(containerHeightRatio)

		// run
		log.Debugf("parameter: %s", parameter.JSONString())
		style.NewStyleProcessor(style.StyleLogoMelon, parameter).Run()
	},
}

func init() {

	flags := LogoMelonCmd.Flags()
	flags.StringP("input", "i", "input", "specify input folder")
	flags.StringP("output", "o", "output", "specify output folder")
	flags.IntP("width", "w", 100, "specify border width")
	flags.StringP("color", "c", "white", "specify border color")
	flags.IntP("container-height", "", 300, "bottom logo container height")
	flags.Float64P("container-height-ratio", "", 0.12, "bottom logo container height, image based height")
	flags.BoolP("without-subtitle", "", false, "without subtitle")
}
