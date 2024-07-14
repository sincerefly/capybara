package exif

type UniversalParser struct {
	meta Meta
}

func NewUniversalParser(meta Meta) *UniversalParser {
	return &UniversalParser{meta: meta}
}

func (p *UniversalParser) Make() (string, error) {
	return p.meta.GetString(TagMake)
}

func (p *UniversalParser) MakeSafe() string {
	return p.meta.GetStringSafe(TagMake)
}

func (p *UniversalParser) Model() (string, error) {
	return p.meta.GetString(TagModel)
}

func (p *UniversalParser) ModelSafe() string {
	return p.meta.GetStringSafe(TagModel)
}

func (p *UniversalParser) FocalLengthIn35mmFormat() (string, error) {
	return p.meta.GetString(TagFocalLengthIn35mmFormat)
}

func (p *UniversalParser) FocalLengthIn35mmFormatSafe() string {
	value, err := p.FocalLengthIn35mmFormat()
	if err != nil {
		return ""
	}
	return value
}

func (p *UniversalParser) Aperture() (string, error) {
	return p.meta.GetString(TagAperture)
}

func (p *UniversalParser) ApertureSafe() string {
	return p.meta.GetStringSafe(TagAperture)
}

func (p *UniversalParser) ShutterSpeed() (string, error) {
	return p.meta.GetString(TagShutterSpeed)
}

func (p *UniversalParser) ShutterSpeedSafe() string {
	return p.meta.GetStringSafe(TagShutterSpeed)
}

func (p *UniversalParser) ISO() (string, error) {
	return p.meta.GetString(TagISO)
}

func (p *UniversalParser) ISOSafe() string {
	return p.meta.GetStringSafe(TagISO)
}

func (p *UniversalParser) LensModel() (string, error) {
	return p.meta.GetString(TagLensModel)
}

func (p *UniversalParser) LensModelSafe() string {
	return p.meta.GetStringSafe(TagLensModel)
}

func (p *UniversalParser) CreateDate() (string, error) {
	return p.meta.GetString(TagCreateDate)
}

func (p *UniversalParser) CreateDateSafe() string {
	return p.meta.GetStringSafe(TagCreateDate)
}
