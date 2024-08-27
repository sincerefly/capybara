package size

type Size struct {
	Width  int
	Height int
}

type Pair struct {
	Source      Size
	Destination Size
}

func NewSizePair(srcWidth, srcHeight, dstWidth, dstHeight int) Pair {
	return Pair{
		Source:      Size{Width: srcWidth, Height: srcHeight},
		Destination: Size{Width: dstWidth, Height: dstHeight},
	}
}

func (p *Pair) SrcWidth() int {
	return p.Source.Width
}

func (p *Pair) SrcHeight() int {
	return p.Source.Height
}

func (p *Pair) DstWidth() int {
	return p.Destination.Width
}

func (p *Pair) SetDstWidth(dstWidth int) {
	p.Destination.Width = dstWidth
}

func (p *Pair) DstHeight() int {
	return p.Destination.Height
}

func (p *Pair) SetDstHeight(dstHeight int) {
	p.Destination.Height = dstHeight
}

type FloatSize struct {
	Width  float64
	Height float64
}
