package resources

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNikonLogoTmpFsPath(t *testing.T) {
	_, err := CreateTemporaryLogoFile("NIKON")
	assert.Nil(t, err)
}
