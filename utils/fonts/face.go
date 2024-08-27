package fonts

import (
	"fmt"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/sincerefly/capybara/resources"
	"github.com/sincerefly/capybara/utils/fsutil"
	"golang.org/x/image/font"
)

// LoadFontFace Load face by ttf font path, with font size
func LoadFontFace(fontPath string, fontSize float64, fontSpecified bool) (font.Face, error) {

	var fontBytes []byte
	var ok bool

	if fontSpecified {
		exist, err := fsutil.Exists(fontPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load specified font: %v, path: %s", err, fontPath)
		}
		if !exist {
			return nil, fmt.Errorf("failed to load specified font: not found, path: %s", fontPath)
		}

		content, err := os.ReadFile(fontPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load specified font: %v, path: %s", err, fontPath)
		}
		fontBytes = content
		ok = true

	}

	if !ok {
		content, err := resources.F.ReadFile(fontPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load built-in font: %v, path: %s", err, fontPath)
		}
		fontBytes = content
	}

	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: fontSize,
		// Hinting: font.HintingFull,
	})
	return face, nil
}
