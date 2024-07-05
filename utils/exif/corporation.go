package exif

import "strings"

type Corporation string

const (
	Sony     Corporation = "SONY"
	Nikon    Corporation = "NIKON"
	Fujifilm Corporation = "FUJIFILM"
	Apple    Corporation = "APPLE"
)

func MakeToCorporation(makeStr string) Corporation {

	upMakeStr := strings.ToUpper(makeStr)
	if strings.Contains(upMakeStr, string(Sony)) {
		return Sony
	} else if strings.Contains(upMakeStr, string(Nikon)) {
		return Nikon
	} else if strings.Contains(upMakeStr, string(Fujifilm)) {
		return Fujifilm
	} else if strings.Contains(upMakeStr, string(Apple)) {
		return Apple
	}
	return ""
}
