package colorizer

import (
	"fmt"
	"golang.org/x/image/colornames"
	"image/color"
	"strings"
)

func htmlNameToColor(name string) (color.RGBA, error) {

	var col color.RGBA
	switch strings.ToLower(name) {

	// Red HTML Color Names
	case "indianred":
		col = color.RGBA{R: 205, G: 92, B: 92, A: 255}
	case "lightcoral":
		col = color.RGBA{R: 240, G: 128, B: 128, A: 255}
	case "salmon":
		col = color.RGBA{R: 250, G: 128, B: 114, A: 255}
	case "darksalmon":
		col = color.RGBA{R: 233, G: 150, B: 122, A: 255}
	case "lightsalmon":
		col = color.RGBA{R: 255, G: 160, B: 122, A: 255}
	case "crimson":
		col = color.RGBA{R: 220, G: 20, B: 60, A: 255}
	case "red":
		col = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	case "firebrick":
		col = color.RGBA{R: 178, G: 34, B: 34, A: 255}
	case "darkred":
		col = color.RGBA{R: 139, G: 0, B: 0, A: 255}

	case "white":
		col = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	case "black":
		col = color.RGBA{R: 0, G: 0, B: 0, A: 255}

	case "test":
		col = colornames.White // https://stackoverflow.com/questions/64612950/function-to-convert-webcolor-html-colour-name-to-hex-in-go-darkorange-ff8c

		// TODO: https://htmlcolorcodes.com/color-names/
	default:
		return color.RGBA{}, fmt.Errorf("invalid html color name: %s", name)
	}

	return col, nil
}
