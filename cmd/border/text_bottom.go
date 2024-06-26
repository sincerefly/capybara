package border

import (
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/service/border"
	"github.com/sincerefly/capybara/service/border/styles"
	"github.com/sincerefly/capybara/utils/colorizer"
	"github.com/spf13/cobra"
)

var TextBottomCmd = &cobra.Command{
	Use:   "text_bottom",
	Short: "Style: Footer text, with photo exif",
	Run: func(cmd *cobra.Command, args []string) {

		tbParameter := &styles.TextBottomParameter{}

		input, _ := cmd.Flags().GetString("input")
		tbParameter.SetInput(input)

		output, _ := cmd.Flags().GetString("output")
		tbParameter.SetOutput(output)

		// width param
		width, _ := cmd.Flags().GetInt("width")
		if borderWidth, fixed := border.FixedBorderWidth(width); fixed {
			log.Warn("border width fixed with %d", borderWidth)
		} else {
			tbParameter.SetBorderWidth(borderWidth)
		}

		// color param
		colorStr, _ := cmd.Flags().GetString("color")
		col, err := colorizer.ToColor(colorStr)
		if err != nil {
			log.Fatal(err)
		}
		tbParameter.SetBorderColor(col)

		// bottom container height
		containerHeight, _ := cmd.Flags().GetInt("container-height")
		tbParameter.SetBottomContainerHeight(containerHeight)

		withoutSubtitle, _ := cmd.Flags().GetBool("without-subtitle")
		tbParameter.SetWithoutSubtitle(withoutSubtitle)

		// run
		border.NewStyleProcessor(border.StyleTextBottom, tbParameter).Run()
	},
}

func init() {

	flags := TextBottomCmd.Flags()
	flags.StringP("input", "i", "input", "specify input folder")
	flags.StringP("output", "o", "output", "specify output folder")
	flags.IntP("width", "w", 100, "specify border width")
	flags.StringP("color", "c", "white", "specify border color")
	flags.IntP("container-height", "", 300, "bottom text container height")
	flags.BoolP("without-subtitle", "", false, "without sub-title")
	flags.BoolP("debug", "d", false, "print details log")
}
