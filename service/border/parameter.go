package border

import (
	"github.com/sincerefly/capybara/utils"
	"image/color"
	"path/filepath"
)

type Style string

const (
	StyleSimple     Style = "simple"
	StyleTextBottom Style = "text_bottom"
)

type GeneralParameter struct {
	style       Style
	input       string
	output      string
	borderWidth int
	borderColor color.Color
}

func NewGeneralParameter() GeneralParameter {
	return GeneralParameter{}
}

func (p *GeneralParameter) SetStyle(style Style) *GeneralParameter {
	p.style = style
	return p
}

func (p *GeneralParameter) SetDefaultStyle() *GeneralParameter {
	p.style = StyleSimple
	return p
}

func (p *GeneralParameter) GetStyle() Style {
	return p.style
}

func (p *GeneralParameter) SetInput(inputDir string) *GeneralParameter {
	p.input = inputDir
	return p
}

func (p *GeneralParameter) SetDefaultInput() *GeneralParameter {
	dir, err := utils.ExecutableDir()
	if err != nil {
		return nil
	}
	p.input = filepath.Join(dir, "input")
	return p
}

func (p *GeneralParameter) GetInput() string {
	return p.input
}

func (p *GeneralParameter) SetOutput(outputDir string) *GeneralParameter {
	p.output = outputDir
	return p
}

func (p *GeneralParameter) SetDefaultOutput() *GeneralParameter {
	dir, err := utils.ExecutableDir()
	if err != nil {
		return nil
	}
	p.output = filepath.Join(dir, "output")
	return p
}

func (p *GeneralParameter) GetOutput() string {
	return p.output
}

func (p *GeneralParameter) SetBorderWidth(borderWidth int) *GeneralParameter {
	p.borderWidth = borderWidth
	return p
}

func (p *GeneralParameter) GetBorderWidth() int {
	return p.borderWidth
}

func (p *GeneralParameter) SetBorderColor(color color.Color) *GeneralParameter {
	p.borderColor = color
	return p
}

func (p *GeneralParameter) GetBorderColor() color.Color {
	return p.borderColor
}

type Parameter interface {
	GetInput() string
	GetOutput() string
}
