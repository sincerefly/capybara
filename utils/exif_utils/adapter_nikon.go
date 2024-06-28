package exif_utils

import (
	"fmt"
	"github.com/sincerefly/capybara/structure/tagname"
	"strconv"
	"strings"
)

type NikonParser struct {
	meta ExifMeta
}

func NewNikonParser(meta ExifMeta) *NikonParser {
	return &NikonParser{meta: meta}
}

func (p *NikonParser) Make() (string, error) {
	return p.meta.GetString(tagname.Make)
}

func (p *NikonParser) MakeSafe() string {
	return p.meta.GetStringSafe(tagname.Make)
}

func (p *NikonParser) Model() (string, error) {
	return p.meta.GetString(tagname.Model)
}

func (p *NikonParser) ModelSafe() string {
	return p.meta.GetStringSafe(tagname.Model)
}

func (p *NikonParser) FocalLengthIn35mmFormat() (string, error) {
	v, err := p.meta.GetString(tagname.FocalLengthIn35mmFormat)
	if err == nil {
		return v, nil
	}

	// focal * scale
	focalStr, err := p.meta.GetString(tagname.FocalLength)
	if err != nil {
		return "", err
	}
	parts := strings.Split(focalStr, " ")

	var focal string
	if len(parts) == 0 {
		return "", fmt.Errorf("%s can't be parserd", tagname.FocalLength)
	}
	focal = parts[0]

	focalFloat, err := strconv.ParseFloat(focal, 64)
	if err != nil {
		return "", err
	}

	scaleFactor35efl, err := p.meta.GetFloat(tagname.ScaleFactor35efl)
	if err != nil {
		return "", err
	}
	focalIn35 := int(focalFloat * scaleFactor35efl)
	focalIn35Str := strconv.Itoa(focalIn35) + " mm"

	return focalIn35Str, nil
}

func (p *NikonParser) FocalLengthIn35mmFormatSafe() string {
	value, err := p.FocalLengthIn35mmFormat()
	if err != nil {
		return ""
	}
	return value
}

func (p *NikonParser) Aperture() (string, error) {
	return p.meta.GetString(tagname.Aperture)
}

func (p *NikonParser) ApertureSafe() string {
	return p.meta.GetStringSafe(tagname.Aperture)
}

func (p *NikonParser) ShutterSpeed() (string, error) {
	return p.meta.GetString(tagname.ShutterSpeed)
}

func (p *NikonParser) ShutterSpeedSafe() string {
	return p.meta.GetStringSafe(tagname.ShutterSpeed)
}

func (p *NikonParser) ISO() (string, error) {
	return p.meta.GetString(tagname.ISO)
}

func (p *NikonParser) ISOSafe() string {
	return p.meta.GetStringSafe(tagname.ISO)
}
