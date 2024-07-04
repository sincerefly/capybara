package exif

import (
	"fmt"
	"os/exec"
)

const DepsExifBin = "exiftool"

func IsExifToolInstalled() (string, error) {
	path, err := exec.LookPath(DepsExifBin)
	if err != nil {
		return "", fmt.Errorf("command '%s' not found: %w", DepsExifBin, err)
	}
	return path, nil
}
