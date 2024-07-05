package styles

import (
	"encoding/json"
	"image/color"
)

type LogoMelonParameter struct {
	input                      string
	output                     string
	borderWidth                int
	borderColor                color.Color
	bottomContainerHeight      int
	isContainerHeightSet       bool
	bottomContainerHeightRatio float64
}

func (p *LogoMelonParameter) JSONString() string {

	resp := map[string]any{
		"input":                      p.Input(),
		"output":                     p.Output(),
		"borderWidth":                p.BorderWidth(),
		"borderColor":                p.BorderColor(),
		"bottomContainerHeight":      p.BottomContainerHeight(),
		"bottomContainerHeightRatio": p.BottomContainerHeightRatio(),
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *LogoMelonParameter) Input() string {
	return p.input
}

func (p *LogoMelonParameter) SetInput(input string) {
	p.input = input
}

func (p *LogoMelonParameter) Output() string {
	return p.output
}

func (p *LogoMelonParameter) SetOutput(output string) {
	p.output = output
}

func (p *LogoMelonParameter) BottomContainerHeight() int {
	return p.bottomContainerHeight
}

func (p *LogoMelonParameter) SetBottomContainerHeight(bottomContainerHeight int) {
	p.bottomContainerHeight = bottomContainerHeight
}

func (p *LogoMelonParameter) IsContainerHeightSet() bool {
	return p.isContainerHeightSet
}

func (p *LogoMelonParameter) SetIsContainerHeightSet(set bool) {
	p.isContainerHeightSet = set
}

func (p *LogoMelonParameter) BottomContainerHeightRatio() float64 {
	return p.bottomContainerHeightRatio
}

func (p *LogoMelonParameter) SetBottomContainerHeightRatio(bottomContainerHeightRatio float64) {
	p.bottomContainerHeightRatio = bottomContainerHeightRatio
}

func (p *LogoMelonParameter) ContainerHeight(srcHeight int) int {
	if p.IsContainerHeightSet() {
		return p.BottomContainerHeight()
	}
	if p.BottomContainerHeightRatio() != 0.0 {
		return int(float64(srcHeight) * p.BottomContainerHeightRatio())
	}
	return p.BottomContainerHeight()
}

func (p *LogoMelonParameter) BorderWidth() int {
	return p.borderWidth
}

func (p *LogoMelonParameter) SetBorderWidth(borderWidth int) {
	p.borderWidth = borderWidth
}

func (p *LogoMelonParameter) BorderColor() color.Color {
	return p.borderColor
}

func (p *LogoMelonParameter) SetBorderColor(color color.Color) {
	p.borderColor = color
}
