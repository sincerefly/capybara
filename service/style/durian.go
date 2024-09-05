package style

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/global"
	"github.com/sincerefly/capybara/resources"
	"github.com/sincerefly/capybara/service/style/styles_common"
	"github.com/sincerefly/capybara/structure/fileitem"
	"github.com/sincerefly/capybara/structure/layout"
	"github.com/sincerefly/capybara/structure/size"
	"github.com/sincerefly/capybara/structure/text"
	"github.com/sincerefly/capybara/utils/exif"
	"github.com/sincerefly/capybara/utils/ggwrapper"
	"golang.org/x/image/colornames"
)

type DurianProcessor struct {
	params  *DurianParameter
	fiStore *fileitem.Store
}

func NewDurianProcessor(params *DurianParameter, fiStore *fileitem.Store) *DurianProcessor {
	return &DurianProcessor{
		params:  params,
		fiStore: fiStore,
	}
}

func (s *DurianProcessor) Run() error {
	if s.fiStore == nil {
		return nil
	}

	// parser exif meta data
	newStore, err := styles_common.SupplementaryMetaToStore(s.fiStore)
	if err != nil {
		return err
	}

	if global.ParamNoParallelism {
		fileitem.LoopExecutor(newStore, s.runner)
	} else {
		fileitem.PoolExecutor(newStore, s.runner)
	}
	return nil
}

func (s *DurianProcessor) runner(fi fileitem.FileItem) error {

	srcImageKey := fi.GetSourceKey()
	outImageKey := fi.GetTargetKey()
	meta := fi.GetExifMeta()

	middleText := meta.ModelSafe()
	rightText := meta.MakeSafe()

	if rightText == "NIKON CORPORATION" { // shorten nikon make
		rightText = "NIKON"
	}

	borderWidth := s.params.borderWidth

	img, err := imaging.Open(srcImageKey, imaging.AutoOrientation(true))
	if err != nil {
		log.Fatalf("failed to open image %v", err)
	}

	imgSizePair := size.NewSizePair(
		img.Bounds().Dx(),
		img.Bounds().Dy(),
		img.Bounds().Dx()+2*borderWidth,
		img.Bounds().Dy(), // rewrite latter
	)

	// resize image for background
	imgResized := imaging.Resize(img, imgSizePair.DstWidth(), 0, imaging.Lanczos)

	imgSizePair.SetDstHeight(imgResized.Bounds().Dy()) // reset dst height

	// new target image
	dst := imaging.New(imgSizePair.DstWidth(), imgResized.Bounds().Dy(), color.RGBA{})

	backgroundImg := imaging.Blur(imgResized, 80) // 50
	dst = imaging.Paste(dst, backgroundImg, image.Pt(0, 0))

	// paste the original image onto a new background
	roundedImg := ggwrapper.ApplyRoundedCorners(img, 150) // 50 是圆角半径，可以根据需要调整

	y := (imgSizePair.DstHeight() - imgSizePair.SrcHeight()) / 3
	dst = imaging.Overlay(dst, roundedImg, image.Pt(borderWidth, y), 1.0)

	// create draw context
	dc := gg.NewContextForImage(dst)

	// draw title and subtitle
	titleSize, err := s.drawTitle(dc, imgSizePair, middleText, rightText)
	if err != nil {
		log.Fatalf("failed to draw title %v", err)
		return err
	}

	if !s.params.WithoutSubtitle() {
		subtitle := s.subtitle(meta)
		err = s.drawSubtitle(dc, imgSizePair, titleSize, subtitle)
		if err != nil {
			log.Fatalf("failed to draw sub-title %v", err)
			return err
		}
	}

	err = imaging.Save(dc.Image(), outImageKey)
	if err != nil {
		log.Fatalf("failed to save image %v", err)
		return err
	}

	log.Infof("with text_bottom saved to %s", outImageKey)
	return nil
}

func (s *DurianProcessor) fixedFontSize(imgSizePair size.Pair) float64 {
	fixedSize := float64((imgSizePair.DstHeight() - imgSizePair.SrcHeight()) / 6)
	if fixedSize > 150 {
		return 150
	}
	return fixedSize
}

func (s *DurianProcessor) drawTitle(dc *gg.Context, imgSizePair size.Pair, middleText, rightText string) (*size.FloatSize, error) {

	const leftText = "Shot on"

	// font size
	fontSize := s.fixedFontSize(imgSizePair)

	leftRt := text.NewRichText(
		leftText,
		resources.AlibabaPuHiTi3LightTTF,
		fontSize,
		color.White,
	)
	middleRt := text.NewRichText(
		middleText,
		resources.AlibabaPuHiTi3BoldTTF,
		fontSize,
		color.White,
	)
	rightRt := text.NewRichText(
		rightText,
		resources.AlibabaPuHiTi3LightTTF,
		fontSize,
		color.White,
	)
	rTexts := []text.RichText{leftRt, middleRt, rightRt}

	newRTexts, textSize := s.textContainerLayout(imgSizePair, nil, rTexts)

	if err := ggwrapper.DrawString(dc, newRTexts); err != nil {
		return nil, err
	}

	return &textSize, nil
}

func (s *DurianProcessor) drawSubtitle(dc *gg.Context, imgSizePair size.Pair, titleSize *size.FloatSize, subtitle string) error {

	fontSize := s.fixedFontSize(imgSizePair)

	richText := text.NewRichText(
		subtitle, // e.g., "70mm f/4.0 1/800s ISO250"
		resources.AlibabaPuHiTi3LightTTF,
		fontSize*0.8,
		colornames.White,
	)
	rTexts := []text.RichText{richText}

	offsetPadding := layout.NewPaddingTop(titleSize.Height * 1.2)

	newRTexts, _ := s.textContainerLayout(imgSizePair, &offsetPadding, rTexts)

	if err := ggwrapper.DrawString(dc, newRTexts); err != nil {
		return err
	}
	return nil
}

func (s *DurianProcessor) subtitle(meta exif.Meta) string {
	focalText := strings.ReplaceAll(meta.FocalLengthIn35mmFormatSafe(), " ", "")
	return fmt.Sprintf("%s f/%s %ss ISO%s", focalText, meta.ApertureSafe(), meta.ShutterSpeedSafe(), meta.ISOSafe())
}

// calculate text container start x,y position
func (s *DurianProcessor) calculateBaseXY(imgSizePair size.Pair, textDim size.FloatSize) layout.Position {

	baseX := float64(imgSizePair.DstWidth()/2) - textDim.Width/2
	baseY := float64((imgSizePair.DstHeight()-imgSizePair.SrcHeight())/3*2+imgSizePair.SrcHeight()) - textDim.Height/2 // - s.fixedHeight()

	return layout.NewPosition(baseX, baseY)
}

func (s *DurianProcessor) textContainerLayout(imgSizePair size.Pair, offsetPadding *layout.Padding,
	rTexts []text.RichText) ([]text.RichText, size.FloatSize) {

	const spacing = " "

	var offsetPaddingLeft, offsetPaddingTop float64
	if offsetPadding != nil {
		offsetPaddingLeft = offsetPadding.PaddingLeft()
		offsetPaddingTop = offsetPadding.PaddingTop()
	}

	paddings := make([]layout.Padding, 0, len(rTexts))

	var paddingLeft float64
	var textContainerWidth, textContainerHeight float64

	// calculate text container size and padding-left info
	for i, rText := range rTexts {
		dc, _ := rText.Context(imgSizePair.DstWidth(), imgSizePair.DstHeight())
		width, height := dc.MeasureString(rText.Text())

		padding := layout.NewPadding(paddingLeft+offsetPaddingLeft, offsetPaddingTop)
		paddings = append(paddings, padding)
		spacingWidth, _ := dc.MeasureString(spacing)
		paddingLeft += width + spacingWidth

		if i != 0 {
			textContainerWidth += spacingWidth
		}

		textContainerWidth += width
		if height > textContainerHeight {
			textContainerHeight = height
		}
	}

	textContainerSize := size.FloatSize{
		Width:  textContainerWidth,
		Height: textContainerHeight,
	}

	basePosition := s.calculateBaseXY(imgSizePair, textContainerSize)

	// append position to rich text
	newRTexts := make([]text.RichText, len(rTexts))
	for i, rText := range rTexts {
		x1 := basePosition.BaseX() + paddings[i].PaddingLeft()
		y1 := basePosition.BaseY() + paddings[i].PaddingTop()
		drawPosition := layout.NewPosition(x1, y1)
		rText.SetPosition(drawPosition)
		newRTexts[i] = rText
	}
	return newRTexts, textContainerSize
}
