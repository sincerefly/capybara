package utils

import (
	"os"
	"path/filepath"
)

func ExecutableDir() (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	executableDir := filepath.Dir(executablePath)
	return executableDir, nil
}
