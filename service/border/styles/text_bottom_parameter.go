package styles

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
		"input":                 p.GetInput(),
		"output":                p.GetOutput(),
		"borderWidth":           p.GetBorderWidth(),
		"borderColor":           p.GetBorderColor(),
		"withoutSubTitle":       p.GetWithoutSubtitle(),
		"bottomContainerHeight": p.GetBottomContainerHeight(),
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *TextBottomParameter) GetInput() string {
	return p.input
}

func (p *TextBottomParameter) SetInput(input string) {
	p.input = input
}

func (p *TextBottomParameter) GetOutput() string {
	return p.output
}

func (p *TextBottomParameter) SetOutput(output string) {
	p.output = output
}

func (p *TextBottomParameter) SetBottomContainerHeight(bottomContainerHeight int) {
	p.bottomContainerHeight = bottomContainerHeight
}

func (p *TextBottomParameter) GetBottomContainerHeight() int {
	return p.bottomContainerHeight
}

func (p *TextBottomParameter) SetBorderWidth(borderWidth int) {
	p.borderWidth = borderWidth
}

func (p *TextBottomParameter) GetBorderWidth() int {
	return p.borderWidth
}

func (p *TextBottomParameter) SetBorderColor(color color.Color) {
	p.borderColor = color
}

func (p *TextBottomParameter) GetBorderColor() color.Color {
	return p.borderColor
}

func (p *TextBottomParameter) SetWithoutSubtitle(withoutSubTitle bool) {
	p.withoutSubTitle = withoutSubTitle
}

func (p *TextBottomParameter) GetWithoutSubtitle() bool {
	return p.withoutSubTitle
}
