package resources

import (
	"embed"
	"errors"
	"github.com/sincerefly/capybara/utils/exif"
	"github.com/sincerefly/capybara/utils/fsutil"
	"os"
	"path/filepath"
	"strings"
)

//go:embed font/* font/*
var F embed.FS

const (
	AlibabaPuHiTi3LightTTF = "font/AlibabaPuHuiTi-3/AlibabaPuHuiTi-3-45-Light.ttf"
	AlibabaPuHiTi3BoldTTF  = "font/AlibabaPuHuiTi-3/AlibabaPuHuiTi-3-85-Bold.ttf"
	RobotBoldTTF           = "font/Robot/Robot-Bold.ttf"
	RobotLightTTF          = "font/Robot/Robot-Light.ttf"
	RobotMediumTTF         = "font/Robot/Robot-Medium.ttf"
	RobotRegularTTF        = "font/Robot/Robot-Regular.ttf"
)

//go:embed logo/*
var LOGO embed.FS

const (
	NikonLogo    = "logo/nikon.png"
	SonyLogo     = "logo/sony.png"
	FujifilmLogo = "logo/fujifilm.png"
	AppleLogo    = "logo/apple.png"
	EmptyLogo    = "logo/empty.png"
)

// CreateTemporaryLogoFile TODO: refactor
func CreateTemporaryLogoFile(makeStr string) (string, error) {

	if makeStr == "" { // check
		return "", errors.New("param 'makeStr' is empty")
	}

	logo := MakeToLogo(makeStr)
	filename := filepath.Join(fsutil.GetTempDir(""), logo)

	has, err := fsutil.Exists(filename) // exist
	if err != nil {
		return "", err
	}
	if has {
		return filename, nil
	}

	b, err := LOGO.ReadFile(logo) // read from embed
	if err != nil {
		return "", err
	}

	exist, err := fsutil.Exists(filepath.Dir(filename))
	if err != nil {
		return "", err
	}
	if !exist {
		fsutil.MkdirAll(filepath.Dir(filename))
	}

	if err = fsutil.WriteFile(filename, b, os.FileMode(0644)); err != nil { // to temp file
		return "", err
	}
	return filename, nil
}

func MakeToLogo(makeStr string) string {

	upMakeStr := strings.ToUpper(makeStr)
	if strings.Contains(upMakeStr, string(exif.SonyCorporation)) {
		return SonyLogo
	} else if strings.Contains(upMakeStr, string(exif.NikonCorporation)) {
		return NikonLogo
	} else if strings.Contains(upMakeStr, string(exif.FujifilmCorporation)) {
		return FujifilmLogo
	} else if strings.Contains(upMakeStr, string(exif.AppleCorporation)) {
		return AppleLogo
	}
	return EmptyLogo
}
