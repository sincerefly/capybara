package utils

import (
	"github.com/spf13/afero"
	"os"
)

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
