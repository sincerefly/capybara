package styles

import "image/color"

type SimpleParameter struct {
	input  string
	output string

	borderWidth int
	borderColor color.Color
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
