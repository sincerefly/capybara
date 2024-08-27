package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNikonLogoTmpFsPath(t *testing.T) {
	_, err := CreateTemporaryLogoFile("NIKON")
	assert.Nil(t, err)
}
