package resources

import (
	"github.com/sincerefly/capybara/base/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNikonLogoTmpFsPath(t *testing.T) {
	path, err := CreateTemporaryLogoFile("NIKON")
	assert.Nil(t, err)
	log.Info(path)
}
