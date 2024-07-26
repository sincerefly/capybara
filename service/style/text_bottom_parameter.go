package style

import (
	"encoding/json"
	"image/color"
)

type TextBottomParameter struct {
	input                 string
	output                string
	borderWidth           int
	borderColor           color.Color
	withoutSubTitle       bool
	bottomContainerHeight int
}

func (p *TextBottomParameter) JSONString() string {

	resp := map[string]any{
		"input":                 p.Input(),
		"output":                p.Output(),
		"borderWidth":           p.BorderWidth(),
		"borderColor":           p.BorderColor(),
		"withoutSubTitle":       p.WithoutSubtitle(),
		"bottomContainerHeight": p.BottomContainerHeight(),
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *TextBottomParameter) Input() string {
	return p.input
}

func (p *TextBottomParameter) SetInput(input string) {
	p.input = input
}

func (p *TextBottomParameter) Output() string {
	return p.output
}

func (p *TextBottomParameter) SetOutput(output string) {
	p.output = output
}

func (p *TextBottomParameter) BottomContainerHeight() int {
	return p.bottomContainerHeight
}

func (p *TextBottomParameter) SetBottomContainerHeight(bottomContainerHeight int) {
	p.bottomContainerHeight = bottomContainerHeight
}

func (p *TextBottomParameter) BorderWidth() int {
	return p.borderWidth
}

func (p *TextBottomParameter) SetBorderWidth(borderWidth int) {
	p.borderWidth = borderWidth
}

func (p *TextBottomParameter) BorderColor() color.Color {
	return p.borderColor
}

func (p *TextBottomParameter) SetBorderColor(color color.Color) {
	p.borderColor = color
}

func (p *TextBottomParameter) WithoutSubtitle() bool {
	return p.withoutSubTitle
}

func (p *TextBottomParameter) SetWithoutSubtitle(withoutSubTitle bool) {
	p.withoutSubTitle = withoutSubTitle
}
