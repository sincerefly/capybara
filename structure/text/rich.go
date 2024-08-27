package text

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/sincerefly/capybara/structure/layout"
	"github.com/sincerefly/capybara/utils/fonts"
)

type RichText struct {
	text          string
	fontPath      string
	fontSpecified bool
	fontSize      float64
	color         color.Color
	position      layout.Position
	anchor        layout.Anchor
}

func NewRichText(text, fontPath string, fontSize float64, color color.Color) RichText {
	return RichText{
		text:     text,
		fontPath: fontPath,
		fontSize: fontSize,
		color:    color,
	}
}

func (rt *RichText) Context(width, height int) (*gg.Context, error) {
	dc := gg.NewContext(width, height)
	dc.SetColor(rt.color)
	face, err := fonts.LoadFontFace(rt.fontPath, rt.fontSize, rt.fontSpecified)
	if err != nil {
		return nil, err
	}
	dc.SetFontFace(face)
	return dc, nil
}

func (rt *RichText) Text() string {
	return rt.text
}

func (rt *RichText) FontPath() string {
	return rt.fontPath
}

func (rt *RichText) FontSpecified() bool {
	return rt.fontSpecified
}

func (rt *RichText) SetFontSpecified() {
	rt.fontSpecified = true
}

func (rt *RichText) FontSize() float64 {
	return rt.fontSize
}

func (rt *RichText) Color() color.Color {
	return rt.color
}

func (rt *RichText) Position() *layout.Position {
	return &rt.position
}

func (rt *RichText) SetPosition(position layout.Position) {
	rt.position = position
}

func (rt *RichText) Anchor() *layout.Anchor {
	return &rt.anchor
}

func (rt *RichText) SetAnchor(anchor layout.Anchor) {
	rt.anchor = anchor
}
