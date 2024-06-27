package styles

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/global"
	"github.com/sincerefly/capybara/resources"
	"github.com/sincerefly/capybara/service/border/styles_common"
	"github.com/sincerefly/capybara/structure"
	"github.com/sincerefly/capybara/structure/layout"
	"github.com/sincerefly/capybara/structure/text_struct"
	"github.com/sincerefly/capybara/utils/exif_utils"
	"github.com/sincerefly/capybara/utils/fileitem"
	"github.com/sincerefly/capybara/utils/gg_utils"
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

	if global.ParamDisableGoroutine {
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
	bottomContainerHeight := s.params.GetBottomContainerHeight()

	img, err := imaging.Open(srcImageKey, imaging.AutoOrientation(true))
	if err != nil {
		log.Fatalf("failed to open image %v", err)
	}

	// image dimensions
	imgDim := structure.ImageDimension{
		SrcWidth:  img.Bounds().Dx(),
		SrcHeight: img.Bounds().Dy(),
		DstWidth:  img.Bounds().Dx() + 2*borderWidth,
		DstHeight: img.Bounds().Dy() + 2*borderWidth + bottomContainerHeight,
	}

	dst := imaging.New(imgDim.DstWidth, imgDim.DstHeight, borderColor)

	// paste the original image onto a new background
	dst = imaging.Paste(dst, img, image.Pt(borderWidth, borderWidth))

	// create draw context
	dc := gg.NewContextForImage(dst)

	// draw title and subtitle
	titleDim, err := s.drawTitle(dc, imgDim, middleText, rightText)
	if err != nil {
		log.Fatalf("failed to draw title %v", err)
		return err
	}

	if !s.params.GetWithoutSubtitle() {
		text := s.subtitleText(meta)
		err = s.drawSubtitle(dc, imgDim, titleDim, text)
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

func (s *TextBottomProcessor) fontSize() float64 {
	size := float64(s.params.bottomContainerHeight / 3)
	if size > 150 {
		return 150
	}
	return size
}

func (s *TextBottomProcessor) fixedHeight() float64 {
	if s.params.GetWithoutSubtitle() {
		return 0
	}
	return float64(s.params.bottomContainerHeight / 10)
}

func (s *TextBottomProcessor) drawTitle(dc *gg.Context, imgDim structure.ImageDimension, middleText, rightText string) (*structure.BaseDimension, error) {

	const leftText = "Shot on"

	// font size
	fontSize := s.fontSize()

	leftRt := text_struct.NewRichText(
		leftText,
		resources.AlibabaPuHiTi3_Light_TTF,
		fontSize,
		color.Black,
	)
	middleRt := text_struct.NewRichText(
		middleText,
		resources.AlibabaPuHiTi3_Bold_TTF,
		fontSize,
		colornames.Red,
	)
	rightRt := text_struct.NewRichText(
		rightText,
		resources.AlibabaPuHiTi3_Bold_TTF,
		fontSize,
		colornames.Black,
	)
	rTexts := []text_struct.RichText{leftRt, middleRt, rightRt}

	newRTexts, textDim := s.textContainerLayout(imgDim, nil, rTexts)

	if err := gg_utils.DrawString(dc, newRTexts); err != nil {
		return nil, err
	}

	return &textDim, nil
}

func (s *TextBottomProcessor) drawSubtitle(dc *gg.Context, imgDim structure.ImageDimension, titleDim *structure.BaseDimension, text string) error {

	fontSize := s.fontSize()

	richText := text_struct.NewRichText(
		text, // e.g., "70mm f/4.0 1/800s ISO250"
		resources.AlibabaPuHiTi3_Light_TTF,
		fontSize*0.8,
		colornames.Gray,
	)
	rTexts := []text_struct.RichText{richText}

	offsetPadding := layout.NewPaddingTop(titleDim.Height * 1.2)

	newRTexts, _ := s.textContainerLayout(imgDim, &offsetPadding, rTexts)

	if err := gg_utils.DrawString(dc, newRTexts); err != nil {
		return err
	}
	return nil
}

func (s *TextBottomProcessor) subtitleText(meta exif_utils.ExifMeta) string {
	focalText := strings.Replace(meta.FocalLengthIn35mmFormatSafe(), " ", "", -1)
	return fmt.Sprintf("%s f/%s %ss ISO%s", focalText, meta.ApertureSafe(), meta.ShutterSpeedSafe(), meta.ISOSafe())
}

// calculate text container start x,y position
func (s *TextBottomProcessor) calculateBaseXY(imgDim structure.ImageDimension, textDim structure.BaseDimension) layout.Position {
	borderWidth := s.params.borderWidth
	bottomContainerHeight := s.params.bottomContainerHeight

	baseX := float64(imgDim.DstWidth/2) - textDim.Width/2
	baseY := float64(2*borderWidth+imgDim.SrcHeight+bottomContainerHeight/2) - textDim.Height/2 - s.fixedHeight()

	return layout.NewPosition(baseX, baseY)
}

func (s *TextBottomProcessor) textContainerLayout(imgDim structure.ImageDimension, offsetPadding *layout.Padding,
	rTexts []text_struct.RichText) ([]text_struct.RichText, structure.BaseDimension) {

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
		dc, _ := rText.Context(imgDim.DstWidth, imgDim.DstHeight)
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

	textContainerDim := structure.BaseDimension{
		Width:  textContainerWidth,
		Height: textContainerHeight,
	}

	basePosition := s.calculateBaseXY(imgDim, textContainerDim)

	// append position to rich text
	newRTexts := make([]text_struct.RichText, len(rTexts))
	for i, rText := range rTexts {
		x1 := basePosition.BaseX() + paddings[i].PaddingLeft()
		y1 := basePosition.BaseY() + paddings[i].PaddingTop()
		drawPosition := layout.NewPosition(x1, y1)
		rText.SetPosition(drawPosition)
		newRTexts[i] = rText
	}
	return newRTexts, textContainerDim
}
