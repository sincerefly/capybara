package style

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/constants"
	"github.com/sincerefly/capybara/structure/fileitem"
	"github.com/sincerefly/capybara/utils/fsutil"
)

type Style string

const (
	StyleSimple     Style = "simple"
	StyleTextBottom Style = "text_bottom"
	StyleLogoMelon  Style = "logo_melon"
	StylePineapple  Style = "pineapple"
)

type Parameterizable interface {
	Input() string
	SetInput(input string)
	Output() string
	SetOutput(output string)
}

type StyleProcessor struct {
	style  Style
	params Parameterizable
}

func NewStyleProcessor(style Style, params Parameterizable) *StyleProcessor {
	return &StyleProcessor{
		style:  style,
		params: params,
	}
}

func (s *StyleProcessor) Run() {

	start := time.Now()
	fiStore, err := s.collectInputs()
	if err != nil {
		log.Fatal(err)
	}

	if err := s.prepareOutputDirs(fiStore); err != nil {
		log.Fatal(err)
	}

	switch s.style {
	case StyleSimple:
		params := s.params.(*SimpleParameter)
		err = NewSimpleProcessor(params, fiStore).Run()
	case StyleTextBottom:
		params := s.params.(*TextBottomParameter)
		err = NewTextBottomProcessor(params, fiStore).Run()
	case StyleLogoMelon:
		params := s.params.(*LogoMelonParameter)
		err = NewLogoMelonProcessor(params, fiStore).Run()
	case StylePineapple:
		params := s.params.(*PineappleParameter)
		err = NewPineappleProcessor(params, fiStore).Run()
	}

	if err != nil {
		log.Fatal(err)
	}
	costs := int(time.Since(start).Seconds())
	log.Infof("finished in %ds", costs)
}

func (s *StyleProcessor) SupportExtensions() []string {
	return []string{constants.ImageExtJPG, constants.ImageExtPNG, constants.ImageExtJPEG}
}

// collect input dir images path
func (s *StyleProcessor) collectInputs() (*fileitem.Store, error) {
	store := fileitem.NewFileItemStore()

	input := s.params.Input()
	output := s.params.Output()

	srcImagePaths, err := fsutil.ListFiles(s.params.Input())
	if err != nil {
		return nil, fmt.Errorf("load input folder failed, %v", err)
	}

	for _, srcImagePath := range srcImagePaths {

		filename := filepath.Base(srcImagePath)

		if !slices.Contains(s.SupportExtensions(), strings.ToLower(filepath.Ext(filename))) {
			continue
		}

		relativePath, err := filepath.Rel(input, srcImagePath)
		if err != nil {
			return nil, err
		}

		fi := fileitem.NewFileItem(filename)
		fi.SetInnerPath(filepath.Dir(relativePath))
		fi.SetSourceBase(input)
		fi.SetTargetBase(output)

		store.Add(fi)

	}
	return &store, nil
}

// prepare output dirs
func (s *StyleProcessor) prepareOutputDirs(fiStore *fileitem.Store) error {
	for _, outputDir := range fiStore.GetTargetPaths() {
		if err := fsutil.MkdirAll(outputDir); err != nil {
			return err
		}
	}
	return nil
}
