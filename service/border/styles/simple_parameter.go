package styles

import (
	"encoding/json"
	"image/color"
)

type SimpleParameter struct {
	input  string
	output string

	borderWidth int
	borderColor color.Color
}

func (p *SimpleParameter) JSONString() string {

	resp := map[string]any{
		"input":       p.GetInput(),
		"output":      p.GetOutput(),
		"borderWidth": p.GetBorderWidth(),
		"borderColor": p.GetBorderColor(),
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *SimpleParameter) GetInput() string {
	return p.input
}

func (p *SimpleParameter) SetInput(input string) {
	p.input = input
}

func (p *SimpleParameter) GetOutput() string {
	return p.output
}

func (p *SimpleParameter) SetOutput(output string) {
	p.output = output
}

func (p *SimpleParameter) SetBorderWidth(borderWidth int) {
	p.borderWidth = borderWidth
}

func (p *SimpleParameter) GetBorderWidth() int {
	return p.borderWidth
}

func (p *SimpleParameter) SetBorderColor(color color.Color) {
	p.borderColor = color
}

func (p *SimpleParameter) GetBorderColor() color.Color {
	return p.borderColor
}
