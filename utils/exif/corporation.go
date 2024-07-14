package exif

import "strings"

type Corporation string

const (
	SonyCorporation     Corporation = "SONY"
	NikonCorporation    Corporation = "NIKON"
	FujifilmCorporation Corporation = "FUJIFILM"
	AppleCorporation    Corporation = "APPLE"
)

func MakeToCorporation(makeStr string) Corporation {

	upMakeStr := strings.ToUpper(makeStr)
	if strings.Contains(upMakeStr, string(SonyCorporation)) {
		return SonyCorporation
	} else if strings.Contains(upMakeStr, string(NikonCorporation)) {
		return NikonCorporation
	} else if strings.Contains(upMakeStr, string(FujifilmCorporation)) {
		return FujifilmCorporation
	} else if strings.Contains(upMakeStr, string(AppleCorporation)) {
		return AppleCorporation
	}
	return ""
}
