package fonts

import (
	"path/filepath"
	"testing"

	"github.com/sincerefly/capybara/resources"
	"github.com/stretchr/testify/assert"
)

func TestLoadFontFace(t *testing.T) {
	_, err := LoadFontFace(resources.AlibabaPuHiTi3LightTTF, 8, false)
	assert.Nil(t, err, "")

	_, err = LoadFontFace(filepath.Join(resources.AlibabaPuHiTi3LightTTF, "_not_found"), 8, false)
	assert.NotNil(t, err, "")

	specifiedPath, _ := filepath.Abs(filepath.Join("../..", "resources", resources.AlibabaPuHiTi3LightTTF))
	_, err = LoadFontFace(specifiedPath, 8, true)
	assert.Nil(t, err, "")

	specifiedPath, _ = filepath.Abs(filepath.Join("../..", "resources", resources.AlibabaPuHiTi3LightTTF, "_not_found"))
	_, err = LoadFontFace(specifiedPath, 8, true)
	assert.NotNil(t, err, "")
}
