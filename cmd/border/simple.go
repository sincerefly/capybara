package border

import (
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/service/border"
	"github.com/sincerefly/capybara/service/border/styles"
	"github.com/sincerefly/capybara/utils/colorizer"
	"github.com/spf13/cobra"
)

var SimpleCmd = &cobra.Command{
	Use:   "simple",
	Short: "Style: add a uniform-width border to the image.",
	Run: func(cmd *cobra.Command, args []string) {

		simpleParameter := &styles.SimpleParameter{}

		input, _ := cmd.Flags().GetString("input")
		simpleParameter.SetInput(input)

		output, _ := cmd.Flags().GetString("output")
		simpleParameter.SetOutput(output)

		// width param
		borderWidth, _ := cmd.Flags().GetInt("width")
		borderWidth, fixed := border.FixedBorderWidth(borderWidth)
		if fixed {
			log.Warn("border width fixed with %d", borderWidth)
		}
		simpleParameter.SetBorderWidth(borderWidth)

		// color param
		colorStr, _ := cmd.Flags().GetString("color")
		col, err := colorizer.ToColor(colorStr)
		if err != nil {
			log.Fatal(err)
		}
		simpleParameter.SetBorderColor(col)

		// run
		border.NewStyleProcessor(border.StyleSimple, simpleParameter).Run()
	},
}

func init() {

	flags := SimpleCmd.Flags()
	flags.StringP("input", "i", "input", "specify input folder")
	flags.StringP("output", "o", "output", "specify output folder")
	flags.IntP("width", "w", 100, "specify border width")
	flags.StringP("color", "c", "white", "specify border color")
	flags.BoolP("debug", "d", false, "print details log")
}
