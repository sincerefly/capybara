package border_common

const (
	MinBorderWidth = 3
	MaxBorderWidth = 2000
)

func FixedBorderWidth(borderWidth int) (width int, fixed bool) {
	if borderWidth > MaxBorderWidth {
		return MaxBorderWidth, true
	} else if borderWidth < MinBorderWidth {
		return MinBorderWidth, true
	}
	return borderWidth, false
}
