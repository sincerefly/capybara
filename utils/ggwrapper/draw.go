package ggwrapper

import (
	"github.com/fogleman/gg"
	"github.com/sincerefly/capybara/structure/text"
	"github.com/sincerefly/capybara/utils/fonts"
)

// DrawString draw string with given position and text with paddings
func DrawString(dc *gg.Context, rTexts []text.RichText) error {
	for _, rt := range rTexts {
		face, err := fonts.LoadFontFace(rt.FontPath(), rt.FontSize(), rt.FontSpecified())
		if err != nil {
			return err
		}
		dc.SetFontFace(face)
		dc.SetColor(rt.Color())
		dc.DrawString(rt.Text(), rt.Position().BaseX(), rt.Position().BaseY())
	}
	return nil
}

func DrawStringAnchored(dc *gg.Context, rTexts []text.RichText) error {
	for _, rt := range rTexts {
		face, err := fonts.LoadFontFace(rt.FontPath(), rt.FontSize(), rt.FontSpecified())
		if err != nil {
			return err
		}
		dc.SetFontFace(face)
		dc.SetColor(rt.Color())
		dc.DrawStringAnchored(rt.Text(), rt.Position().BaseX(), rt.Position().BaseY(), rt.Anchor().Ax(), rt.Anchor().Ay())
	}
	return nil
}
