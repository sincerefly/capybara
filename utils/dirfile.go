package utils

import (
	"github.com/spf13/afero"
	"os"
	"path/filepath"
)

var appFs afero.Fs

func MkdirAll(path string) error {

	if appFs == nil {
		appFs = afero.NewOsFs()
	}
	fs := afero.NewOsFs()
	return fs.MkdirAll(path, 0755)
}

// GetAllFiles all files
func GetAllFiles(fs afero.Fs, root string) ([]string, error) {
	var files []string

	err := afero.Walk(fs, root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil
}

func ExecutableDir() (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	executableDir := filepath.Dir(executablePath)
	return executableDir, nil
}
