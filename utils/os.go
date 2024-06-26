package utils

import (
	"github.com/spf13/afero"
)

var appFs afero.Fs

func MkdirAll(path string) error {

	if appFs == nil {
		appFs = afero.NewOsFs()
	}

	fs := afero.NewOsFs()
	return fs.MkdirAll(path, 0755)
}
