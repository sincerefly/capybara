package exif_utils

import (
	"github.com/sincerefly/capybara/structure/tagname"
)

type UniversalParser struct {
	meta ExifMeta
}

func NewUniversalParser(meta ExifMeta) *UniversalParser {
	return &UniversalParser{meta: meta}
}

func (p *UniversalParser) Make() (string, error) {
	return p.meta.GetString(tagname.Make)
}

func (p *UniversalParser) MakeSafe() string {
	return p.meta.GetStringSafe(tagname.Make)
}

func (p *UniversalParser) Model() (string, error) {
	return p.meta.GetString(tagname.Model)
}

func (p *UniversalParser) ModelSafe() string {
	return p.meta.GetStringSafe(tagname.Model)
}

func (p *UniversalParser) FocalLengthIn35mmFormat() (string, error) {
	return p.meta.GetString(tagname.FocalLengthIn35mmFormat)
}

func (p *UniversalParser) FocalLengthIn35mmFormatSafe() string {
	value, err := p.FocalLengthIn35mmFormat()
	if err != nil {
		return ""
	}
	return value
}

func (p *UniversalParser) Aperture() (string, error) {
	return p.meta.GetString(tagname.Aperture)
}

func (p *UniversalParser) ApertureSafe() string {
	return p.meta.GetStringSafe(tagname.Aperture)
}

func (p *UniversalParser) ShutterSpeed() (string, error) {
	return p.meta.GetString(tagname.ShutterSpeed)
}

func (p *UniversalParser) ShutterSpeedSafe() string {
	return p.meta.GetStringSafe(tagname.ShutterSpeed)
}

func (p *UniversalParser) ISO() (string, error) {
	return p.meta.GetString(tagname.ISO)
}

func (p *UniversalParser) ISOSafe() string {
	return p.meta.GetStringSafe(tagname.ISO)
}
