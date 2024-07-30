package style

import (
	"encoding/json"
	"image/color"
)

type PineappleParameter struct {
	input     string
	output    string
	fontColor color.Color
}

func (p *PineappleParameter) JSONString() string {

	resp := map[string]any{
		"input":     p.Input(),
		"output":    p.Output(),
		"fontColor": p.FontColor(),
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *PineappleParameter) Input() string {
	return p.input
}

func (p *PineappleParameter) SetInput(input string) {
	p.input = input
}

func (p *PineappleParameter) Output() string {
	return p.output
}

func (p *PineappleParameter) SetOutput(output string) {
	p.output = output
}

func (p *PineappleParameter) FontColor() color.Color {
	return p.fontColor
}

func (p *PineappleParameter) SetFontColor(color color.Color) {
	p.fontColor = color
}
