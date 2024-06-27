package styles

import (
	"github.com/disintegration/imaging"
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/utils/fileitem"
	"image"
)

type SimpleProcessor struct {
	params  *SimpleParameter
	fiStore *fileitem.Store
}

func NewSimpleProcessor(params *SimpleParameter, fiStore *fileitem.Store) *SimpleProcessor {
	return &SimpleProcessor{
		params:  params,
		fiStore: fiStore,
	}
}

func (s *SimpleProcessor) Run() error {
	if s.fiStore == nil {
		return nil
	}

	for _, fi := range s.fiStore.GetItems() {
		s.runner(fi.GetSourceKey(), fi.GetTargetKey())
	}
	return nil
}

func (s *SimpleProcessor) runner(srcImageKey, outputImageKey string) {

	borderWidth := s.params.GetBorderWidth()
	borderColor := s.params.GetBorderColor()

	img, err := imaging.Open(srcImageKey, imaging.AutoOrientation(true))
	if err != nil {
		log.Fatalf("failed to open image %v", err)
	}

	// src image dimension
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	// dst image dimension
	newWidth := width + 2*borderWidth
	newHeight := height + 2*borderWidth

	dst := imaging.New(newWidth, newHeight, borderColor) // dst image

	dst = imaging.Paste(dst, img, image.Pt(borderWidth, borderWidth)) // paste image to dst

	// save image
	err = imaging.Save(dst, outputImageKey)
	if err != nil {
		log.Fatalf("failed to save image %v", err)
	}
	log.Infof("image with simple border saved to %s", outputImageKey)
}