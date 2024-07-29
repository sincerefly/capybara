package style

import (
	"fmt"
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
	"image"
	"image/color"
	"time"
)

type PineappleProcessor struct {
	params  *PineappleParameter
	fiStore *fileitem.Store
}

func NewPineappleProcessor(params *PineappleParameter, fiStore *fileitem.Store) *PineappleProcessor {
	return &PineappleProcessor{
		params:  params,
		fiStore: fiStore,
	}
}

func (s *PineappleProcessor) Run() error {
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

func (s *PineappleProcessor) runner(fi fileitem.FileItem) error {

	srcImageKey := fi.GetSourceKey()
	outImageKey := fi.GetTargetKey()
	meta := fi.GetExifMeta()

	fontColor := s.params.FontColor()

	img, err := imaging.Open(srcImageKey, imaging.AutoOrientation(true))
	if err != nil {
		log.Fatalf("failed to open image %v", err)
	}

	imgSize := size.Size{
		Width:  img.Bounds().Dx(),
		Height: img.Bounds().Dy(),
	}

	dst := imaging.New(imgSize.Width, imgSize.Height, color.White)

	// paste the original image onto a new background
	dst = imaging.Paste(dst, img, image.Pt(0, 0))

	// create draw context
	dc := gg.NewContextForImage(dst)

	// draw text
	err = s.drawShotTime(dc, imgSize, meta, fontColor)
	if err != nil {
		log.Fatalf("failed to draw title %v", err)
		return err
	}

	err = imaging.Save(dc.Image(), outImageKey)
	if err != nil {
		log.Fatalf("failed to save image %v", err)
		return err
	}

	log.Infof("with pineapple style saved to %s", outImageKey)
	return nil
}

func (s *PineappleProcessor) fixedFontSize(size size.Size) float64 {
	if size.Height > size.Width {
		return float64(size.Height) / 25 // Experience numbers
	}
	return float64(size.Height) / 22
}

func (s *PineappleProcessor) drawShotTime(dc *gg.Context, imgSize size.Size, meta exif.Meta, fontColor color.Color) error {

	// font size
	fontSize := s.fixedFontSize(imgSize)

	createDate, err := s.metaDateToDateStampFormat(meta.CreateDateSafe())
	if err != nil {
		log.Fatalf("convert datetime failed, %v", err)
	}
	//fontColor, _ := colorizer.ToColor("rgba(255, 140, 0, 230)")

	dateTimeRt := text.NewRichText(
		createDate,
		resources.Digital7MonoTTF,
		fontSize,
		fontColor,
	)
	rtDc, _ := dateTimeRt.Context(imgSize.Width, imgSize.Height)
	fontWidth, _ := rtDc.MeasureString(dateTimeRt.Text())

	s.fillLeftTopPos(&dateTimeRt, imgSize, fontSize, fontWidth)

	rTexts := []text.RichText{dateTimeRt}

	if err := ggwrapper.DrawString(dc, rTexts); err != nil {
		return err
	}
	return nil
}

func (s *PineappleProcessor) fillLeftTopPos(rText *text.RichText, imgSize size.Size, fontSize float64, fontWidth float64) {

	x := float64(imgSize.Width) - fontWidth - fontSize
	y := float64(imgSize.Height) - fontSize + 50
	drawPosition := layout.NewPosition(x, y)
	rText.SetPosition(drawPosition)
}

// Output Format: '24 07 29 21:22
func (s *PineappleProcessor) metaDateToDateStampFormat(input string) (string, error) {
	// Define the layout of the input time string
	const inputLayout = "2006:01:02 15:04:05"

	// Parse the input string to time.Time format
	t, err := time.Parse(inputLayout, input)
	if err != nil {
		return "", err
	}

	// Define the desired output format
	year := fmt.Sprintf("'%02d", t.Year()%100)
	month := fmt.Sprintf("%02d", t.Month())
	day := fmt.Sprintf("%02d", t.Day())
	hours := fmt.Sprintf("%02d", t.Hour())
	minutes := fmt.Sprintf("%02d", t.Minute())

	// Concatenate to form the output string
	output := fmt.Sprintf("%s %s %s %s:%s", year, month, day, hours, minutes)
	return output, nil
}
