package border

const (
	MinBorderWidth = 3
	MaxBorderWidth = 2000
)

func FixedBorderWidth(borderWidth int) (width int, fixed bool) {
	if borderWidth > 2000 {
		return MaxBorderWidth, true
	} else if borderWidth < 3 {
		return MinBorderWidth, true
	}
	return borderWidth, false
}
