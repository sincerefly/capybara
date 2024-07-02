package structure

type IntSize struct {
	Width  int
	Height int
}

type IntSizePair struct {
	Source      IntSize
	Destination IntSize
}

func NewIntSizePair(srcWidth, srcHeight, dstWidth, dstHeight int) IntSizePair {
	return IntSizePair{
		Source:      IntSize{Width: srcWidth, Height: srcHeight},
		Destination: IntSize{Width: dstWidth, Height: dstHeight},
	}
}

func (p *IntSizePair) SrcWidth() int {
	return p.Source.Width
}

func (p *IntSizePair) SrcHeight() int {
	return p.Source.Height
}

func (p *IntSizePair) DstWidth() int {
	return p.Destination.Width
}

func (p *IntSizePair) DstHeight() int {
	return p.Destination.Height
}

type FloatSize struct {
	Width  float64
	Height float64
}
