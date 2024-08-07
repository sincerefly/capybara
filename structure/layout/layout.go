package layout

type Position struct {
	x float64
	y float64
}

func NewPosition(x, y float64) Position {
	return Position{x: x, y: y}
}

func (p *Position) BaseX() float64 {
	return p.x
}

func (p *Position) BaseY() float64 {
	return p.y
}

type Anchor struct {
	ax float64
	ay float64
}

func NewAnchor(ax, ay float64) Anchor {
	return Anchor{ax: ax, ay: ay}
}

func (p *Anchor) Ax() float64 {
	return p.ax
}

func (p *Anchor) Ay() float64 {
	return p.ay
}

type Padding struct {
	left float64
	top  float64
}

func NewPadding(left, top float64) Padding {
	return Padding{
		left: left,
		top:  top,
	}
}

func NewPaddingLeft(left float64) Padding {
	return Padding{
		left: left,
	}
}

func NewPaddingTop(top float64) Padding {
	return Padding{
		top: top,
	}
}

func (p *Padding) PaddingLeft() float64 {
	return p.left
}

func (p *Padding) PaddingTop() float64 {
	return p.top
}
