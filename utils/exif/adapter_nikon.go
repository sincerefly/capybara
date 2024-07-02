package exif

import (
	"fmt"
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
	return p.meta.GetString(TagMake)
}

func (p *NikonParser) MakeSafe() string {
	return p.meta.GetStringSafe(TagMake)
}

func (p *NikonParser) Model() (string, error) {
	return p.meta.GetString(TagModel)
}

func (p *NikonParser) ModelSafe() string {
	return p.meta.GetStringSafe(TagModel)
}

func (p *NikonParser) FocalLengthIn35mmFormat() (string, error) {
	v, err := p.meta.GetString(TagFocalLengthIn35mmFormat)
	if err == nil {
		return v, nil
	}

	// focal * scale
	focalStr, err := p.meta.GetString(TagFocalLength)
	if err != nil {
		return "", err
	}
	parts := strings.Split(focalStr, " ")

	var focal string
	if len(parts) == 0 {
		return "", fmt.Errorf("%s can't be parserd", TagFocalLength)
	}
	focal = parts[0]

	focalFloat, err := strconv.ParseFloat(focal, 64)
	if err != nil {
		return "", err
	}

	scaleFactor35efl, err := p.meta.GetFloat(TagScaleFactor35efl)
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
	return p.meta.GetString(TagAperture)
}

func (p *NikonParser) ApertureSafe() string {
	return p.meta.GetStringSafe(TagAperture)
}

func (p *NikonParser) ShutterSpeed() (string, error) {
	return p.meta.GetString(TagShutterSpeed)
}

func (p *NikonParser) ShutterSpeedSafe() string {
	return p.meta.GetStringSafe(TagShutterSpeed)
}

func (p *NikonParser) ISO() (string, error) {
	return p.meta.GetString(TagISO)
}

func (p *NikonParser) ISOSafe() string {
	return p.meta.GetStringSafe(TagISO)
}
