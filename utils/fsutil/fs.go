package fsutil

import (
	"github.com/spf13/afero"
	"os"
	"path/filepath"
)

var appFs afero.Fs

func init() {
	appFs = afero.NewOsFs()
}

func MkdirAll(path string) error {
	fs := afero.NewOsFs()
	return fs.MkdirAll(path, 0755)
}

func ListFiles(root string) ([]string, error) {
	var paths []string
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	}
	err := afero.Walk(appFs, root, walkFn)
	if err != nil {
		return nil, err
	}
	return paths, nil
}

func ExecutableDir() (string, error) {

	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	executableDir := filepath.Dir(executablePath)
	return executableDir, nil
}
