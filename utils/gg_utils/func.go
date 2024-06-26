package gg_utils

import (
	"github.com/fogleman/gg"
	"github.com/sincerefly/capybara/structure/text_struct"
	"github.com/sincerefly/capybara/utils"
)

// DrawString draw string with given position and text with paddings
func DrawString(dc *gg.Context, rTexts []text_struct.RichText) error {
	for _, rt := range rTexts {
		face, err := utils.LoadFontFace(rt.FontPath(), rt.FontSize(), rt.FontSpecified())
		if err != nil {
			return err
		}
		dc.SetFontFace(face)
		dc.SetColor(rt.Color())
		dc.DrawString(rt.Text(), rt.Position().BaseX(), rt.Position().BaseY())
	}
	return nil
}
