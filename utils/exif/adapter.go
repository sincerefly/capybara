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
	LensModel() (string, error)
	LensModelSafe() string
	CreateDate() (string, error)
	CreateDateSafe() string
}

func makeAdapter(meta Meta) Adapter {
	switch MakeToCorporation(meta.GetStringSafe(TagMake)) {
	case NikonCorporation:
		return NewNikonParser(meta)
	default:
		return NewUniversalParser(meta)
	}
}

func (m *Meta) Make() (string, error) {
	return makeAdapter(*m).Make()
}

func (m *Meta) MakeSafe() string {
	return makeAdapter(*m).MakeSafe()
}

func (m *Meta) Model() (string, error) {
	return makeAdapter(*m).Model()
}

func (m *Meta) ModelSafe() string {
	return makeAdapter(*m).ModelSafe()
}

func (m *Meta) FocalLengthIn35mmFormat() (string, error) {
	return makeAdapter(*m).FocalLengthIn35mmFormat()
}

func (m *Meta) FocalLengthIn35mmFormatSafe() string {
	return makeAdapter(*m).FocalLengthIn35mmFormatSafe()
}

func (m *Meta) Aperture() (string, error) {
	return makeAdapter(*m).Aperture()
}

func (m *Meta) ApertureSafe() string {
	return makeAdapter(*m).ApertureSafe()
}

func (m *Meta) ShutterSpeed() (string, error) {
	return makeAdapter(*m).ShutterSpeed()
}

func (m *Meta) ShutterSpeedSafe() string {
	return makeAdapter(*m).ShutterSpeedSafe()
}

func (m *Meta) ISO() (string, error) {
	return makeAdapter(*m).ISO()
}

func (m *Meta) ISOSafe() string {
	return makeAdapter(*m).ISOSafe()
}

func (m *Meta) LensModel() (string, error) {
	return makeAdapter(*m).LensModel()
}

func (m *Meta) LensModelSafe() string {
	return makeAdapter(*m).LensModelSafe()
}

func (m *Meta) CreateDate() (string, error) {
	return makeAdapter(*m).CreateDate()
}

func (m *Meta) CreateDateSafe() string {
	return makeAdapter(*m).CreateDateSafe()
}
