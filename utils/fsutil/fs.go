package fsutil

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
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

func Exists(path string) (bool, error) {
	return afero.Exists(appFs, path)
}

func GetTempDir(path string) string {
	return afero.GetTempDir(appFs, path)
}

func TempFile(dir string, pattern string) (f afero.File, err error) {
	return afero.TempFile(appFs, dir, pattern)
}

func WriteFile(filename string, data []byte, perm os.FileMode) error {
	return afero.WriteFile(appFs, filename, data, perm)
}

func ExecutableDir() (string, error) {

	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	executableDir := filepath.Dir(executablePath)
	return executableDir, nil
}
