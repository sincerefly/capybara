package styles

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/global"
	"github.com/sincerefly/capybara/resources"
	"github.com/sincerefly/capybara/service/border/styles_common"
	"github.com/sincerefly/capybara/structure/fileitem"
	"github.com/sincerefly/capybara/structure/layout"
	"github.com/sincerefly/capybara/structure/size"
	"github.com/sincerefly/capybara/structure/text"
	"github.com/sincerefly/capybara/utils/exif"
	"github.com/sincerefly/capybara/utils/ggwrapper"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"strings"
)

type TextBottomProcessor struct {
	params  *TextBottomParameter
	fiStore *fileitem.Store
}

func NewTextBottomProcessor(params *TextBottomParameter, fiStore *fileitem.Store) *TextBottomProcessor {
	return &TextBottomProcessor{
		params:  params,
		fiStore: fiStore,
	}
}

func (s *TextBottomProcessor) Run() error {
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

func (s *TextBottomProcessor) runner(fi fileitem.FileItem) error {

	srcImageKey := fi.GetSourceKey()
	outImageKey := fi.GetTargetKey()
	meta := fi.GetExifMeta()

	middleText := meta.ModelSafe()
	rightText := meta.MakeSafe()

	borderWidth := s.params.borderWidth
	borderColor := s.params.borderColor
	bottomContainerHeight := s.params.BottomContainerHeight()

	img, err := imaging.Open(srcImageKey, imaging.AutoOrientation(true))
	if err != nil {
		log.Fatalf("failed to open image %v", err)
	}

	imgSizePair := size.NewSizePair(
		img.Bounds().Dx(),
		img.Bounds().Dy(),
		img.Bounds().Dx()+2*borderWidth,
		img.Bounds().Dy()+2*borderWidth+bottomContainerHeight,
	)

	dst := imaging.New(imgSizePair.DstWidth(), imgSizePair.DstHeight(), borderColor)

	// paste the original image onto a new background
	dst = imaging.Paste(dst, img, image.Pt(borderWidth, borderWidth))

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

	log.Infof("image with simple border saved to %s", outImageKey)
	return nil
}

func (s *TextBottomProcessor) fixedFontSize() float64 {
	fixedSize := float64(s.params.bottomContainerHeight / 3)
	if fixedSize > 150 {
		return 150
	}
	return fixedSize
}

func (s *TextBottomProcessor) fixedHeight() float64 {
	if s.params.WithoutSubtitle() {
		return 0
	}
	return float64(s.params.bottomContainerHeight / 10)
}

func (s *TextBottomProcessor) drawTitle(dc *gg.Context, imgSizePair size.Pair, middleText, rightText string) (*size.FloatSize, error) {

	const leftText = "Shot on"

	// font size
	fontSize := s.fixedFontSize()

	leftRt := text.NewRichText(
		leftText,
		resources.AlibabaPuHiTi3LightTTF,
		fontSize,
		color.Black,
	)
	middleRt := text.NewRichText(
		middleText,
		resources.AlibabaPuHiTi3BoldTTF,
		fontSize,
		colornames.Red,
	)
	rightRt := text.NewRichText(
		rightText,
		resources.AlibabaPuHiTi3BoldTTF,
		fontSize,
		colornames.Black,
	)
	rTexts := []text.RichText{leftRt, middleRt, rightRt}

	newRTexts, textSize := s.textContainerLayout(imgSizePair, nil, rTexts)

	if err := ggwrapper.DrawString(dc, newRTexts); err != nil {
		return nil, err
	}

	return &textSize, nil
}

func (s *TextBottomProcessor) drawSubtitle(dc *gg.Context, imgSizePair size.Pair, titleSize *size.FloatSize, subtitle string) error {

	fontSize := s.fixedFontSize()

	richText := text.NewRichText(
		subtitle, // e.g., "70mm f/4.0 1/800s ISO250"
		resources.AlibabaPuHiTi3LightTTF,
		fontSize*0.8,
		colornames.Gray,
	)
	rTexts := []text.RichText{richText}

	offsetPadding := layout.NewPaddingTop(titleSize.Height * 1.2)

	newRTexts, _ := s.textContainerLayout(imgSizePair, &offsetPadding, rTexts)

	if err := ggwrapper.DrawString(dc, newRTexts); err != nil {
		return err
	}
	return nil
}

func (s *TextBottomProcessor) subtitle(meta exif.Meta) string {
	focalText := strings.ReplaceAll(meta.FocalLengthIn35mmFormatSafe(), " ", "")
	return fmt.Sprintf("%s f/%s %ss ISO%s", focalText, meta.ApertureSafe(), meta.ShutterSpeedSafe(), meta.ISOSafe())
}

// calculate text container start x,y position
func (s *TextBottomProcessor) calculateBaseXY(imgSizePair size.Pair, textDim size.FloatSize) layout.Position {
	borderWidth := s.params.borderWidth
	bottomContainerHeight := s.params.bottomContainerHeight

	baseX := float64(imgSizePair.DstWidth()/2) - textDim.Width/2
	baseY := float64(2*borderWidth+imgSizePair.SrcHeight()+bottomContainerHeight/2) - textDim.Height/2 - s.fixedHeight()

	return layout.NewPosition(baseX, baseY)
}

func (s *TextBottomProcessor) textContainerLayout(imgSizePair size.Pair, offsetPadding *layout.Padding,
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
