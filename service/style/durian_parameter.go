package style

import (
	"encoding/json"
)

type DurianParameter struct {
	input           string
	output          string
	borderWidth     int
	withoutSubTitle bool
}

func (p *DurianParameter) JSONString() string {

	resp := map[string]any{
		"input":           p.Input(),
		"output":          p.Output(),
		"borderWidth":     p.BorderWidth(),
		"withoutSubTitle": p.WithoutSubtitle(),
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *DurianParameter) Input() string {
	return p.input
}

func (p *DurianParameter) SetInput(input string) {
	p.input = input
}

func (p *DurianParameter) Output() string {
	return p.output
}

func (p *DurianParameter) SetOutput(output string) {
	p.output = output
}

func (p *DurianParameter) BorderWidth() int {
	return p.borderWidth
}

func (p *DurianParameter) SetBorderWidth(borderWidth int) {
	p.borderWidth = borderWidth
}

func (p *DurianParameter) WithoutSubtitle() bool {
	return p.withoutSubTitle
}

func (p *DurianParameter) SetWithoutSubtitle(withoutSubTitle bool) {
	p.withoutSubTitle = withoutSubTitle
}
