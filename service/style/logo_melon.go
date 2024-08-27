package style

import (
	"fmt"
	"image"
	"image/color"
	"strings"
	"time"

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
)

type LogoMelonProcessor struct {
	params  *LogoMelonParameter
	fiStore *fileitem.Store
}

func NewLogoMelonProcessor(params *LogoMelonParameter, fiStore *fileitem.Store) *LogoMelonProcessor {
	return &LogoMelonProcessor{
		params:  params,
		fiStore: fiStore,
	}
}

func (s *LogoMelonProcessor) Run() error {
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

func (s *LogoMelonProcessor) runner(fi fileitem.FileItem) error {

	srcImageKey := fi.GetSourceKey()
	outImageKey := fi.GetTargetKey()
	meta := fi.GetExifMeta()

	borderWidth := s.params.borderWidth
	borderColor := s.params.borderColor

	img, err := imaging.Open(srcImageKey, imaging.AutoOrientation(true))
	if err != nil {
		log.Fatalf("failed to open image %v", err)
	}
	bottomContainerHeight := s.params.ContainerHeight(img.Bounds().Dy())

	imgSizePair := size.NewSizePair(
		img.Bounds().Dx(),
		img.Bounds().Dy(),
		img.Bounds().Dx()+borderWidth*2,
		img.Bounds().Dy()+borderWidth+bottomContainerHeight,
	)

	dst := imaging.New(imgSizePair.DstWidth(), imgSizePair.DstHeight(), borderColor)

	// paste the original image onto a new background
	dst = imaging.Paste(dst, img, image.Pt(borderWidth, borderWidth))

	logoImg := s.loadLogo(meta.MakeSafe(), bottomContainerHeight) // prepare logo image
	dst = imaging.Overlay(dst, logoImg, image.Pt(borderWidth, imgSizePair.SrcHeight()+borderWidth), 1.0)

	// create draw context
	dc := gg.NewContextForImage(dst)

	// draw text
	err = s.drawText(dc, imgSizePair, logoImg, meta)
	if err != nil {
		log.Fatalf("failed to draw title %v", err)
		return err
	}

	err = imaging.Save(dc.Image(), outImageKey)
	if err != nil {
		log.Fatalf("failed to save image %v", err)
		return err
	}

	log.Infof("with melon style saved to %s", outImageKey)
	return nil
}

func (s *LogoMelonProcessor) loadLogo(mateMake string, bottomContainerHeight int) image.Image {

	// logo fs temp path
	fsLogo, err := resources.CreateTemporaryLogoFile(mateMake)
	if err != nil {
		log.Fatalf("failed to prepare temporary logo file %v", err)
	}

	logoImg, err := imaging.Open(fsLogo, imaging.AutoOrientation(true))
	if err != nil {
		log.Fatalf("failed to open image %v", err)
	}
	desiredHeight := bottomContainerHeight // Change this value to your desired height

	// Resize the logoImg to the desired height while maintaining the aspect ratio
	return imaging.Resize(logoImg, 0, desiredHeight, imaging.Lanczos)
}

func (s *LogoMelonProcessor) fixedFontSize(srcHeight int) float64 {
	fixedSize := float64(s.params.ContainerHeight(srcHeight) / 5)
	if fixedSize > 150 {
		return 150
	}
	return fixedSize
}

func (s *LogoMelonProcessor) drawText(dc *gg.Context, imgPair size.Pair, logoImg image.Image, meta exif.Meta) error {

	// font size
	fontSize := s.fixedFontSize(imgPair.SrcHeight())
	logoWidth := logoImg.Bounds().Dx()

	var lens = meta.LensModelSafe()
	if meta.LensModelSafe() == "" {
		lens = "Built-in Lens " + meta.FocalLengthIn35mmFormatSafe()
	}
	leftTopRt := text.NewRichText(
		lens,
		resources.AlibabaPuHiTi3BoldTTF,
		fontSize,
		color.Black,
	)
	s.fillLeftTopPos(&leftTopRt, imgPair, logoWidth, fontSize)

	leftBottomRt := text.NewRichText(
		meta.ModelSafe(),
		resources.AlibabaPuHiTi3LightTTF,
		fontSize,
		color.Black,
	)
	s.fillLeftBottomPos(&leftBottomRt, imgPair, logoWidth)

	rightTopRt := text.NewRichText(
		s.imgDescriptionText(meta),
		resources.AlibabaPuHiTi3BoldTTF,
		fontSize,
		color.Black,
	)
	rtDc, _ := rightTopRt.Context(imgPair.DstWidth(), imgPair.DstHeight())
	fontWidth, _ := rtDc.MeasureString(rightTopRt.Text())
	rightTopX := s.fillRightTopPos(&rightTopRt, imgPair, fontSize, fontWidth)

	createDate, err := convertDateTime(meta.CreateDateSafe())
	if err != nil {
		log.Fatalf("convert datetime failed, %v", err)
	}
	rightBottomRt := text.NewRichText(
		createDate,
		resources.AlibabaPuHiTi3LightTTF,
		fontSize,
		color.Black,
	)
	rtDc, _ = rightTopRt.Context(imgPair.DstWidth(), imgPair.DstHeight())
	fontWidth, _ = rtDc.MeasureString(rightBottomRt.Text())
	s.fillRightBottomPos(&rightBottomRt, imgPair, fontWidth, rightTopX)

	rTexts := []text.RichText{leftTopRt, leftBottomRt, rightTopRt, rightBottomRt}

	if err := ggwrapper.DrawString(dc, rTexts); err != nil {
		return err
	}
	return nil
}

func (s *LogoMelonProcessor) fillLeftTopPos(rText *text.RichText, imgPair size.Pair, logoWidth int, fontSize float64) {

	border := s.params.BorderWidth()
	srcHeight := imgPair.SrcHeight()

	x := float64(logoWidth + border + 50)
	y := float64(srcHeight+border+s.params.ContainerHeight(srcHeight)/4) + fontSize
	drawPosition := layout.NewPosition(x, y)
	rText.SetPosition(drawPosition)
}

func (s *LogoMelonProcessor) fillLeftBottomPos(rText *text.RichText, imgPair size.Pair, logoWidth int) {

	border := s.params.BorderWidth()
	srcHeight := imgPair.SrcHeight()

	x := float64(logoWidth + border + 50)
	y := float64(srcHeight + border + s.params.ContainerHeight(srcHeight)/4*3)
	drawPosition := layout.NewPosition(x, y)
	rText.SetPosition(drawPosition)
}

func (s *LogoMelonProcessor) fillRightTopPos(rText *text.RichText, imgPair size.Pair, fontSize, fontWidth float64) float64 {

	border := s.params.BorderWidth()
	srcHeight := imgPair.SrcHeight()

	x := float64(imgPair.DstWidth()-border-int(fontWidth)) - 50
	y := float64(srcHeight+border+s.params.ContainerHeight(srcHeight)/4) + fontSize
	drawPosition := layout.NewPosition(x, y)
	rText.SetPosition(drawPosition)
	return x
}

func (s *LogoMelonProcessor) fillRightBottomPos(rText *text.RichText, imgPair size.Pair, fontWidth, rightTopX float64) {

	border := s.params.BorderWidth()
	srcHeight := imgPair.SrcHeight()

	x := float64(imgPair.DstWidth() - border - int(fontWidth))
	if rightTopX > 0 {
		x = rightTopX
	}
	y := float64(srcHeight + border + s.params.ContainerHeight(srcHeight)/4*3)
	drawPosition := layout.NewPosition(x, y)
	rText.SetPosition(drawPosition)
}

func (s *LogoMelonProcessor) imgDescriptionText(meta exif.Meta) string {
	focalText := strings.ReplaceAll(meta.FocalLengthIn35mmFormatSafe(), " ", "")
	return fmt.Sprintf("%s f/%s %ss ISO%s", focalText, meta.ApertureSafe(), meta.ShutterSpeedSafe(), meta.ISOSafe())
}

func convertDateTime(input string) (string, error) {
	// Define the layout of the input time string
	const inputLayout = "2006:01:02 15:04:05"
	// Define the layout for the output time string
	const outputLayout = "2006-01-02 15:04:05"

	// Parse the input string to time.Time format
	t, err := time.Parse(inputLayout, input)
	if err != nil {
		return "", err
	}

	// Format the time to the desired output layout
	output := t.Format(outputLayout)
	return output, nil
}
