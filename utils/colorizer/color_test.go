package colorizer

import (
	"github.com/magiconair/properties/assert"
	"log"
	"testing"
)

// https://htmlcolorcodes.com/color-names/

func TestToColor(t *testing.T) {
	hexColorStr := "#FF69B4"
	color, err := ToColor(hexColorStr)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, rgbaToHex(color), hexColorStr+"FF", "")
}
