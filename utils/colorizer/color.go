package colorizer

import (
	"fmt"
	"golang.org/x/image/colornames"
	"image/color"
	"strconv"
	"strings"
)

func ToColor(colorStr string) (color.RGBA, error) {

	var col color.RGBA
	var err error

	if strings.HasPrefix(colorStr, "#") {
		col, err = hexToColor(colorStr)
	} else if strings.HasPrefix(colorStr, "rgb(") && strings.HasSuffix(colorStr, ")") {
		col, err = rgbToColor(colorStr)
	} else if strings.HasPrefix(colorStr, "rgba(") && strings.HasSuffix(colorStr, ")") {
		col, err = rgbaToColor(colorStr)
	} else {
		//col, err = htmlNameToColor(colorStr)
		value, ok := colornames.Map[strings.ToLower(colorStr)]
		if ok {
			col = value
		} else {
			return color.RGBA{}, fmt.Errorf("invalid color: %s", colorStr)
		}
	}

	if err != nil {
		return color.RGBA{}, fmt.Errorf("parser color failed: %s", err)
	}
	return col, nil
}

func hexToColor(hex string) (color.RGBA, error) {
	hex = hex[1:]
	var rgba color.RGBA

	if len(hex) == 6 {
		hex += "FF" // alpha default 'FF'
	}

	// 解析红、绿、蓝、透明度通道值
	if len(hex) == 8 {
		r, _ := strconv.ParseUint(hex[0:2], 16, 8)
		g, _ := strconv.ParseUint(hex[2:4], 16, 8)
		b, _ := strconv.ParseUint(hex[4:6], 16, 8)
		a, _ := strconv.ParseUint(hex[6:8], 16, 8)

		rgba = color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(a),
		}
	} else {
		return color.RGBA{}, fmt.Errorf("invalid hex color code")
	}
	return rgba, nil
}

func rgbaToColor(rgbaStr string) (color.RGBA, error) {
	rgbaStr = rgbaStr[5 : len(rgbaStr)-1] // remove 'rgba(' and ')'
	parts := strings.Split(rgbaStr, ",")

	if len(parts) != 4 {
		return color.RGBA{}, fmt.Errorf("invalid rgba color code")
	}

	r, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return color.RGBA{}, err
	}

	g, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return color.RGBA{}, err
	}

	b, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		return color.RGBA{}, err
	}

	a, err := strconv.Atoi(strings.TrimSpace(parts[3]))
	if err != nil {
		return color.RGBA{}, err
	}
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}, nil
}

func rgbToColor(rgbStr string) (color.RGBA, error) {
	rgbStr = rgbStr[4 : len(rgbStr)-1] // remove 'rgb(' and ')'
	parts := strings.Split(rgbStr, ",")

	if len(parts) != 3 {
		return color.RGBA{}, fmt.Errorf("invalid rgb color code")
	}

	r, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return color.RGBA{}, err
	}

	g, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return color.RGBA{}, err
	}

	b, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		return color.RGBA{}, err
	}
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}, nil
}

func rgbaToHex(c color.RGBA) string {
	return fmt.Sprintf("#%02X%02X%02X%02X", c.R, c.G, c.B, c.A)
}
