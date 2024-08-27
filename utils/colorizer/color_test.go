package colorizer

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://htmlcolorcodes.com/color-names/

func TestToColor(t *testing.T) {
	hexColorStr := "#FF69B4"
	color, err := ToColor(hexColorStr)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, rgbaToHex(color), hexColorStr+"FF", "")
}

func TestToColor_Named(t *testing.T) {
	colorStr := "red"
	color, err := ToColor(colorStr)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, rgbaToHex(color), "#FF0000FF", "")
}

func TestToColor_hexToColor(t *testing.T) {
	hexColorStr := "#FF69B4"

	color, err := hexToColor(hexColorStr)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, rgbaToHex(color), hexColorStr+"FF", "")
}

// web: https://www.rgbtohex.net/ RGB to HEX Color Converter
func TestToColor_rgbToColor(t *testing.T) {
	rgbColorStr := "rgb(238,130,238)"

	color, err := rgbToColor(rgbColorStr)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, rgbaToHex(color), "#EE82EEFF", "")
}

func TestToColor_rgbaToColor(t *testing.T) {
	rgbaColorStr := "rgba(238,130,238,255)"

	color, err := rgbaToColor(rgbaColorStr)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, rgbaToHex(color), "#EE82EEFF", "")
}
