package fsutil

import (
	"fmt"
	"testing"
)

func TestGetTempDir(t *testing.T) {
	temp := GetTempDir("logo")
	fmt.Println(temp)
}
