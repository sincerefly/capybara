package exif

type Adapter interface {
	Make() (string, error)
	MakeSafe() string
	Model() (string, error)
	ModelSafe() string
	FocalLengthIn35mmFormat() (string, error)
	FocalLengthIn35mmFormatSafe() string
	Aperture() (string, error)
	ApertureSafe() string
	ShutterSpeed() (string, error)
	ShutterSpeedSafe() string
	ISO() (string, error)
	ISOSafe() string
}

func makeAdapter(meta ExifMeta) Adapter {
	switch meta.GetStringSafe(TagMake) {
	case "Nikon":
		return NewNikonParser(meta)
	default:
		return NewUniversalParser(meta)
	}
}

func (m *ExifMeta) Make() (string, error) {
	return makeAdapter(*m).Make()
}

func (m *ExifMeta) MakeSafe() string {
	return makeAdapter(*m).MakeSafe()
}

func (m *ExifMeta) Model() (string, error) {
	return makeAdapter(*m).Model()
}

func (m *ExifMeta) ModelSafe() string {
	return makeAdapter(*m).ModelSafe()
}

func (m *ExifMeta) FocalLengthIn35mmFormat() (string, error) {
	return makeAdapter(*m).FocalLengthIn35mmFormat()
}

func (m *ExifMeta) FocalLengthIn35mmFormatSafe() string {
	return makeAdapter(*m).FocalLengthIn35mmFormatSafe()
}

func (m *ExifMeta) Aperture() (string, error) {
	return makeAdapter(*m).Aperture()
}

func (m *ExifMeta) ApertureSafe() string {
	return makeAdapter(*m).ApertureSafe()
}

func (m *ExifMeta) ShutterSpeed() (string, error) {
	return makeAdapter(*m).ShutterSpeed()
}

func (m *ExifMeta) ShutterSpeedSafe() string {
	return makeAdapter(*m).ShutterSpeedSafe()
}

func (m *ExifMeta) ISO() (string, error) {
	return makeAdapter(*m).ISO()
}

func (m *ExifMeta) ISOSafe() string {
	return makeAdapter(*m).ISOSafe()
}
