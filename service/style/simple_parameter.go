package style

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
		"input":       p.Input(),
		"output":      p.Output(),
		"borderWidth": p.BorderWidth(),
		"borderColor": p.GetBorderColor(),
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *SimpleParameter) Input() string {
	return p.input
}

func (p *SimpleParameter) SetInput(input string) {
	p.input = input
}

func (p *SimpleParameter) Output() string {
	return p.output
}

func (p *SimpleParameter) SetOutput(output string) {
	p.output = output
}

func (p *SimpleParameter) SetBorderWidth(borderWidth int) {
	p.borderWidth = borderWidth
}

func (p *SimpleParameter) BorderWidth() int {
	return p.borderWidth
}

func (p *SimpleParameter) SetBorderColor(color color.Color) {
	p.borderColor = color
}

func (p *SimpleParameter) GetBorderColor() color.Color {
	return p.borderColor
}
